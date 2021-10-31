package command

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/message"
	"github.com/ytake/morse-message/publisher/sub"
)

type MessageSubscriber struct {
	Client *sub.Receiver
}

// Run Kafka Topic Subscriber
func (m *MessageSubscriber) Run(_ *cli.Context) error {
	return m.Client.Subscribe(&message.NoKeyMessagesReceive{})
}
