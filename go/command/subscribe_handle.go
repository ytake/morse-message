package command

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/message"
	"github.com/ytake/morse-message/publisher/sub"
)

type SinglePartitionSubscriber struct {
	Client *sub.Receiver
}

type MultiplePartitionSubscriber struct {
	Client *sub.Receiver
}

// Run Kafka Topic Subscriber
func (m *SinglePartitionSubscriber) Run(_ *cli.Context) error {
	return m.Client.Subscribe(&message.NoKeyMessagesReceive{})
}

// Run Kafka Topic Subscriber
func (m *MultiplePartitionSubscriber) Run(_ *cli.Context) error {
	return m.Client.Subscribe(&message.NoKeyMessagesReceive{})
}
