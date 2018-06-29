package navio

import "context"

type Client struct {
	cc *clientConn
}

// New create new client with target and options set
func New(target string,opts ...ClientOption) (*Client) {
	c:=&Client{
		cc: &clientConn{
			target:target,
		},
	}
	for _, opt := range opts {
		opt(&c.cc.opts)
	}

	return c
}

func (c *Client) Dial () error {
	return c.DialWithContext(context.Background())
}

func (c *Client) DialWithContext (ctx context.Context) (err error) {
	c.cc.ctx, c.cc.cancel = context.WithCancel(ctx)
	defer func() {
		select {
		case <-ctx.Done():
			c.cc.gClient, err = nil, ctx.Err()
		default:
		}
		if err != nil {
			c.cc.close()
		}
	}()
	err =  c.cc.getClientConn()
	return
}


// Close tear down the connection to navio server
func (c *Client) Close () error {
	return c.cc.close()
}

func (c *Client) NewMessage (topic,meta string, body []byte) *Message{
	return newMessage(c.cc,topic,meta,body)
}
func (c *Client) M () *Message{
	return m(c.cc)
}

func (c *Client) NewRequest (id, topic,meta string, body []byte,timeout int) *SendRequest{
	return newRequest(c.cc,id,topic,meta,body,timeout)
}
func (c *Client) R () *SendRequest{
	return r(c.cc)
}

func (c *Client) NewMessageStream(ch chan *Message) error {
	return c.cc.sendMessageStream(ch)
}

func (c *Client) SubscribeToTopic(topic string, ch chan*Message) error {
	return c.cc.subscribeToChannel(topic,"",ch)
}

func (c *Client) SubscribeToTopicWithQueueGroup(topic,group string, ch chan*Message) error {
	return c.cc.subscribeToChannel(topic,group,ch)
}
func (c *Client) SubscribeToRequest(topic string, ch chan *GetRequest) error {
	return c.cc.subscribeToRequests(topic,"",ch)
}

func (c *Client) SubscribeToRequestWithQueueGroup(topic,group string, ch chan*GetRequest) error {
	return c.cc.subscribeToRequests(topic,group,ch)
}