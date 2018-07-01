package navio

type Message struct {
	Topic string
	Meta  string
	Body  []byte
	cc    *clientConn
}

func newMessage(cc *clientConn, topic, meta string, body []byte) *Message {
	return &Message{
		Topic: topic,
		Meta:  meta,
		Body:  body,
		cc:    cc,
	}
}

func m(cc *clientConn) *Message {
	return &Message{
		cc: cc,
	}
}

func (m *Message) SetTopic(value string) *Message {
	m.Topic = value
	return m
}

func (m *Message) SetMetadata(value string) *Message {
	m.Meta = value
	return m
}

func (m *Message) SetBody(value []byte) *Message {
	m.Body = value
	return m
}

func (m *Message) Send() error {
	return m.cc.sendMessage(m, 0)
}

func (m *Message) SendWithTimeout(t int) error {
	return m.cc.sendMessage(m, t)
}
