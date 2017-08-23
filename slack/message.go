package slack

type Message struct {
	conn *Connection
	channel,
	Text string
}

func (m *Message) Respond(answer string) {
	m.conn.rtm.SendMessage(m.conn.rtm.NewOutgoingMessage(answer, m.channel))
}
