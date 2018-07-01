package navio

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	ot "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	pb "github.com/liornabat/navio-go/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"strconv"
)

type clientConn struct {
	ctx     context.Context
	cancel  context.CancelFunc
	target  string
	opts    options
	gClient pb.WarpClient
	gConn   *grpc.ClientConn
	dopts   []grpc.DialOption
}

func createMetaDataContext(headers map[string]string) context.Context {
	h := metadata.New(headers)
	return metadata.NewOutgoingContext(context.Background(), h)
}

func createMetaDataContextWithContext(ctx context.Context, headers map[string]string) context.Context {
	h := metadata.New(headers)
	return metadata.NewOutgoingContext(ctx, h)
}

func (cc *clientConn) optionsApply() error {
	if cc.opts.name == "" {
		cc.opts.name = fmt.Sprintf("navio_clinet_%s", uuid.New().String()[0:5])
	}

	if cc.opts.isSecure {
		var err error
		creds, err := credentials.NewClientTLSFromFile(cc.opts.certFile, cc.opts.keyFile)
		if err != nil {
			return fmt.Errorf("could not load tls cert: %s", err)
		}
		cc.dopts = append(cc.dopts, grpc.WithTransportCredentials(creds))
	} else {
		cc.dopts = append(cc.dopts, grpc.WithInsecure())
	}

	if cc.opts.isTracing {
		cc.dopts = append(cc.dopts, grpc.WithUnaryInterceptor(ot.UnaryClientInterceptor(ot.WithTracer(cc.opts.tracer))))
		cc.dopts = append(cc.dopts, grpc.WithStreamInterceptor(ot.StreamClientInterceptor(ot.WithTracer(cc.opts.tracer))))
	}

	if cc.opts.block {
		cc.dopts = append(cc.dopts, grpc.WithBlock())
	}
	return nil
}

func (cc *clientConn) getClientConn() (err error) {
	if err = cc.optionsApply(); err != nil {
		return
	}
	if cc.gConn, err = grpc.Dial(cc.target, cc.dopts...); err != nil {
		return
	}
	cc.gClient = pb.NewWarpClient(cc.gConn)
	return
}

func (cc *clientConn) close() error {
	defer cc.cancel()
	return cc.gConn.Close()
}

func (cc *clientConn) sendMessage(msg *Message, timeout int) (err error) {
	md := make(map[string]string)
	md["client_tag"] = cc.opts.name
	if timeout > 0 {
		md["timeout"] = strconv.Itoa(timeout)
	}
	ctx := createMetaDataContext(md)
	pbMessage := &pb.Message{
		Channel:  msg.Topic,
		Metadata: msg.Meta,
		Body:     msg.Body,
	}
	_, err = cc.gClient.SendMessage(ctx, pbMessage)
	return
}

func (cc *clientConn) subscribeToChannel(topic, group string, ch chan *Message) error {
	md := make(map[string]string)
	md["client_tag"] = cc.opts.name
	ctx := createMetaDataContext(md)
	subRequest := &pb.SubscribeRequest{
		Channel: topic,
		Group:   group,
	}
	sub, err := cc.gClient.SubscribeToChannel(ctx, subRequest)
	if err != nil {
		return err
	}
	go func() {
		for {
			msg, err := sub.Recv()
			if err != nil {
				close(ch)
				return
			}
			ch <- &Message{
				Topic: msg.Channel,
				Meta:  msg.Metadata,
				Body:  msg.Body,
			}
		}

	}()
	return nil
}

func (cc *clientConn) sendMessageStream(ch chan *Message) error {
	md := make(map[string]string)
	md["client_tag"] = cc.opts.name
	ctx := createMetaDataContext(md)
	sender, err := cc.gClient.SendMessageStream(ctx)
	if err != nil {
		return err
	}
	go func() {
		for {
			msg, closed := <-ch
			if closed {
				return
			}
			pbMessage := &pb.Message{
				Channel:  msg.Topic,
				Metadata: msg.Meta,
				Body:     msg.Body,
			}
			err := sender.Send(pbMessage)
			if err != nil {
				return
			}
		}
	}()
	return nil
}
func (cc *clientConn) sendRequest(ctx context.Context, req *SendRequest) (*GetResponse, error) {
	md := make(map[string]string)
	md["client_tag"] = cc.opts.name
	ctx = createMetaDataContextWithContext(ctx, md)
	pbRequest := &pb.Request{
		ID:       req.ID,
		Channel:  req.Topic,
		Metadata: req.Meta,
		Timeout:  int32(req.Timeout),
		Body:     req.Body,
		CacheKey: req.CacheKey,
		CacheTTL: int32(req.CacheTTL),
	}
	pbResponse, err := cc.gClient.SendRequest(ctx, pbRequest)
	if err != nil {
		return nil, err
	}
	response := &GetResponse{
		RequestID: pbResponse.RequestID,
		Meta:      pbResponse.Metadata,
		Body:      pbResponse.Body,
		CacheHit:  pbResponse.CacheHit,
	}
	return response, nil
}
func (cc *clientConn) sendResponse(ctx context.Context, req *GetRequest) error {
	md := make(map[string]string)
	md["client_tag"] = cc.opts.name
	ctx = createMetaDataContextWithContext(ctx, md)
	pbResponse := &pb.Response{
		RequestID:    req.ID,
		Body:         req.Body,
		Metadata:     req.Meta,
		ReplyChannel: req.replyTopic,
	}
	_, err := cc.gClient.SendResponse(ctx, pbResponse)
	if err != nil {
		return err
	}
	return nil
}

func (cc *clientConn) subscribeToRequests(topic, group string, ch chan *GetRequest) error {
	md := make(map[string]string)
	md["client_tag"] = cc.opts.name
	md["channel"] = topic
	md["group"] = group
	ctx := createMetaDataContext(md)
	sub, err := cc.gClient.RequestResponseStream(ctx)
	if err != nil {
		return err
	}
	go func() {
		for {
			req, err := sub.Recv()
			if err != nil {
				close(ch)
				return
			}
			if req.ReplyChannel == "" {
				continue
			}
			ch <- &GetRequest{
				ID:         req.ID,
				Meta:       req.Metadata,
				Body:       req.Body,
				replyTopic: req.ReplyChannel,
			}
		}
	}()
	return nil
}
