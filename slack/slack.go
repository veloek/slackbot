package slack

import (
	nlopes "github.com/nlopes/slack"
)

func Connect(token string) *Connection {
	client := nlopes.New(token)
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	conn := &Connection{
		client:   client,
		rtm:      rtm,
		Messages: make(chan Message, 50),
	}
	go conn.readEvents()

	return conn
}

type Connection struct {
	client   *nlopes.Client
	rtm      *nlopes.RTM
	botInfo  *nlopes.Info
	Messages chan Message
}

func (c *Connection) Dispose() error {
	err := c.rtm.Disconnect()
	return err
}

func (c *Connection) readEvents() {
	for msg := range c.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *nlopes.MessageEvent:
			if ev.SubType == "" && isDirectMessage(ev) {
				c.Messages <- Message{conn: c, channel: ev.Channel, Text: ev.Text}
			}
		default:
			// Not handling other kinds of events for now
		}
	}
}

func isDirectMessage(ev *nlopes.MessageEvent) bool {
	return ev.Channel[:1] == "D"
}
