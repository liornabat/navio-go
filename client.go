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

