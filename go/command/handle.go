package command

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/kafka"
)

type MessagePublisher struct {
	Client *kafka.Client
}

func (m *MessagePublisher) Run(_ *cli.Context) error {

	return nil
}
