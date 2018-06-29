package navio

import (
	"fmt"
	"context"
)

type SendRequest struct {
	ID string
	Topic string
	Meta string
	Body []byte
	Timeout int
	CacheKey string
	CacheTTL int
	cc *clientConn
}

type GetResponse struct {
	RequestID string
	Meta string
	Body []byte
	CacheHit bool
}

func newRequest(cc *clientConn,id, topic,meta string, body []byte,timeout int) *SendRequest {
	return &SendRequest{
		ID:id,
		Topic:topic,
		Meta:meta,
		Body:body,
		Timeout:timeout,
		cc:cc,
	}
}

func r(cc *clientConn) *SendRequest {
	return &SendRequest{
		cc:cc,
	}
}

func (r *SendRequest) SetID(value string )*SendRequest {
	r.ID=value
	return r
}
func (r *SendRequest) SetTimeout(value int )*SendRequest {
	r.Timeout=value
	return r
}
func (r *SendRequest) SetTopic(value string )*SendRequest {
	r.Topic=value
	return r
}

func (r *SendRequest) SetMetadata(value string )*SendRequest {
	r.Meta=value
	return r
}

func (r *SendRequest) SetBody(value []byte )*SendRequest {
	r.Body=value
	return r
}

func (r *SendRequest) SetCache(key string, ttl int )*SendRequest {
	r.CacheKey=key
	r.CacheTTL=ttl
	return r
}

func (r *SendRequest) validate()error{
	if r.Topic=="" {
		return fmt.Errorf("invalid topic, cannot be empty")
	}
	if len(r.Body)==0 {
		return fmt.Errorf("invalid body, cannot be empty")
	}
	if r.Timeout==0 {
		return fmt.Errorf("invalid timeout, cannot be 0")
	}
	return nil
}

func (r *SendRequest) Send (ctx context.Context) (*GetResponse, error) {
	if err:=r.validate();err!=nil {
		return nil,err
	}
	return r.cc.sendRequest(ctx,r)
}


type GetRequest struct {
	ID string
	Meta string
	Body []byte
	replyTopic string
	cc *clientConn
}

func (gr *GetRequest) SendResponse(ctx context.Context, meta string,body []byte) error {
	if err:=gr.validate();err!=nil {
		return err
	}
	return gr.cc.sendResponse(ctx,gr)
}
func (gr *GetRequest) validate()error{
	if gr.replyTopic=="" {
		return fmt.Errorf("invalid reply topic, cannot be empty")
	}
	return nil
}