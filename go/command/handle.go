package command

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/kafka"
	"github.com/ytake/morse-message/publisher/message"
	"google.golang.org/protobuf/proto"
)

type MessagePublisher struct {
	Client *kafka.Client
}

func (m *MessagePublisher) Run(_ *cli.Context) error {
	messages, err := message.NoKeyMessages()
	if err != nil {
		return err
	}
	for _, v := range messages {
		o, err := proto.Marshal(v)
		if err != nil {
			return err
		}
		m.Client.Publish(o)
	}
	m.Client.Close()
	return nil
}
