package console

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/command"
	"github.com/ytake/morse-message/publisher/message"
)

type (
	SinglePartitionPublisher struct {
		Command command.UserAction
	}
	NoKeyPartitionPublisher struct {
		Command command.UserAction
	}
)

func (m *SinglePartitionPublisher) Run(_ *cli.Context) error {
	messages, err := message.NoKeyMessages()
	if err != nil {
		return err
	}
	return m.Command.UserRegistrationForNoKey(messages)
}

func (m *NoKeyPartitionPublisher) Run(_ *cli.Context) error {
	messages, err := message.NoKeyMessages()
	if err != nil {
		return err
	}
	return m.Command.UserRegistrationForNoKey(messages)
}
