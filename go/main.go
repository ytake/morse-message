package main

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/command"
	"github.com/ytake/morse-message/publisher/config"
	"github.com/ytake/morse-message/publisher/log"
	"github.com/ytake/morse-message/publisher/pub"
	"github.com/ytake/morse-message/publisher/sub"
	"go.uber.org/zap"
	"os"
)

func main() {
	c := config.New()
	l := log.NewLogger()
	defer l.Flush()
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "message:nokey_publish",
				Aliases: []string{"m:np"},
				Usage:   "to Kafka",
				Action: func(context *cli.Context) error {
					pc, err := pub.NewProducer(c.Kafka.KafkaBootstrapServers())
					if err != nil {
						l.Error("pub producer error", zap.Error(err))
						return err
					}
					cmp := &command.MessagePublisher{Client: pub.NewNoKeyClient(c.Kafka.UserActionTopic(), pc)}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:nokey_subscriber",
				Aliases: []string{"m:ns"},
				Usage:   "from Kafka",
				Action: func(context *cli.Context) error {
					sc, err := sub.NewConsumer(c.Kafka.KafkaBootstrapServers(), "nokey_subscriber_sample")
					if err != nil {
						l.Error("subscriber error", zap.Error(err))
						return err
					}
					cmp := &command.MessageSubscriber{Client: sub.NewNoKeyClient(c.Kafka.UserActionTopic(), sc)}
					return cmp.Run(context)
				},
			},
		},
	}
	app.Name = "SplitBrainMessage (do not use for your application)"
	err := app.Run(os.Args)
	if err != nil {
		l.ServerFatal("console error", zap.Error(err))
	}
}
