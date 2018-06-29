package navio

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"fmt"
	ot "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	pb "github.com/liornabat/navio-go/protobuf"
	"github.com/google/uuid"
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


func (cc *clientConn) optionsApply () error {
	if cc.opts.name=="" {
		cc.opts.name=fmt.Sprintf("navio_clinet_%s",uuid.New().String()[0:5])
	}

	if cc.opts.isSecure {
		var err error
		creds, err := credentials.NewClientTLSFromFile(cc.opts.certFile, cc.opts.keyFile)
		if err != nil {
			return fmt.Errorf("could not load tls cert: %s", err)
		}
		cc.dopts=append(cc.dopts,grpc.WithTransportCredentials(creds))
	}else {
		cc.dopts=append(cc.dopts,grpc.WithInsecure())
	}

	if cc.opts.isTracing {
		cc.dopts=append(cc.dopts,grpc.WithUnaryInterceptor(ot.UnaryClientInterceptor(ot.WithTracer(cc.opts.tracer))))
		cc.dopts=append(cc.dopts,grpc.WithStreamInterceptor(ot.StreamClientInterceptor(ot.WithTracer(cc.opts.tracer))))
	}

	if cc.opts.block {
		cc.dopts=append(cc.dopts,grpc.WithBlock())
	}
	return nil
}

func (cc *clientConn) getClientConn() (err error) {
	if err=cc.optionsApply();err!=nil {
		return
	}
	if cc.gConn, err = grpc.Dial(cc.target,cc.dopts...);err!=nil {
		return
	}
	cc.gClient = pb.NewWarpClient(cc.gConn)
	return
}

func (cc *clientConn) close () error {
	defer cc.cancel()
	return cc.gConn.Close()
}

func (cc *clientConn) sendMessage (msg *Message,timeout int) (err error) {
	md:=make(map[string]string)
	md["client_tag"] =cc.opts.name
	if timeout>0 {
		md["timeout"]=strconv.Itoa(timeout)
	}
	ctx := createMetaDataContext(md)
	pbMessage:=&pb.Message{
		Channel:msg.Topic,
		Metadata:msg.Meta,
		Body:msg.Body,
	}
	_,err = cc.gClient.SendMessage(ctx,pbMessage)
	return
}





















