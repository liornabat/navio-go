package tests

import (
	"context"
	"fmt"
	pb "github.com/liornabat/navio-go/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
	"google.golang.org/grpc/metadata"
	"strings"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/google/uuid"
)

type Server struct {
	port        string
	grpcServer  *grpc.Server
	mu          sync.RWMutex
	messagesMap map[string]map[string]chan*pb.Message
	requestsMap map[string]map[string]chan*pb.Request
	responseMap map[string]chan*pb.Response
}
func getFromMetadata(ctx context.Context, key string) (string, bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	if len(headers[key]) == 0 {
		return "", false
	}

	return strings.Join(headers[key], ""), true
}

func NewServer(port, cert, key string) (s *Server, err error) {
	s = &Server{
		port:        port,
		messagesMap: make(map[string]map[string]chan *pb.Message),
		requestsMap: make(map[string]map[string]chan *pb.Request),
		responseMap: make (map[string]chan *pb.Response),
	}
	serverOptions := []grpc.ServerOption{}
	if cert != "" {
		creds, err := credentials.NewServerTLSFromFile(cert, key)
		if err != nil {
			return nil, fmt.Errorf("could not load TLS keys: %s", err)
		}
		serverOptions = append(serverOptions, grpc.Creds(creds))
	}
	s.grpcServer = grpc.NewServer(serverOptions...)
	reflection.Register(s.grpcServer)
	pb.RegisterWarpServer(s.grpcServer, s)
	return s, err
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		panic(err.Error())
	}
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			panic(err.Error())
		}
	}()
}

func (s *Server) sendMessage (msg *pb.Message)  {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, topicMap := range s.messagesMap[msg.Channel] {
		topicMap<-msg
	}
}

func (s *Server) SendMessage(ctx context.Context, msg *pb.Message) (empty *pb.Empty, err error) {
	s.sendMessage(msg)
	return &pb.Empty{}, nil
}

func (s *Server) SendMessageStream(stream pb.Warp_SendMessageStreamServer) error {
	for {
		msg,err:=stream.Recv()
		if err != nil {
			return nil
		}
		s.sendMessage(msg)
	}
	return nil
}

func (s *Server) SubscribeToChannel(req *pb.SubscribeRequest, stream pb.Warp_SubscribeToChannelServer) (err error) {
	s.mu.Lock()
	_,ok:=s.messagesMap[req.Channel]
	if !ok {
		s.messagesMap[req.Channel]=make(map[string]chan*pb.Message)
		s.messagesMap[req.Channel][req.Group]=make (chan *pb.Message,1)
	} else {
		_,ok:=s.messagesMap[req.Channel][req.Group]
		if !ok {
			s.messagesMap[req.Channel][req.Group]=make (chan *pb.Message,1)
		}
	}
	s.mu.Unlock()
	for {
		select {
		case msg:=<-s.messagesMap[req.Channel][req.Group]:
			err:=stream.Send(msg)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func (s *Server) SendRequest(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	s.mu.Lock()
	replyChId:=uuid.New().String()
	s.responseMap[replyChId]=make(chan *pb.Response,1)
	s.mu.Unlock()
	req.ReplyChannel=replyChId
	
	return &pb.Response{}, nil
}

func (s *Server) SendResponse(ctx context.Context, res *pb.Response) (empty *pb.Empty, err error) {
	return &pb.Empty{}, nil
}

func (s *Server) RequestResponseStream(stream pb.Warp_RequestResponseStreamServer) error {
	s.mu.Lock()
	channel, ok := getFromMetadata(stream.Context(), "channel")
	if !ok {
		return status.Error(codes.Internal, "no channel was set")
	}
	group, _ := getFromMetadata(stream.Context(), "group")

	_,ok=s.requestsMap[channel]
	if !ok {
		s.requestsMap[channel]=make(map[string]chan*pb.Request)
		s.requestsMap[channel][group]=make (chan *pb.Request,1)
	} else {
		_,ok:=s.requestsMap[channel][group]
		if !ok {
			s.requestsMap[channel][group]=make (chan *pb.Request,1)
		}
	}
	s.mu.Unlock()
	for {
		select {
		case req:=<-s.requestsMap[channel][group]:
			err:=stream.Send(req)
			if err != nil {
				return nil
			}
		}
	}
	return nil


}
