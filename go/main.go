package main

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/command"
	"github.com/ytake/morse-message/publisher/config"
	"github.com/ytake/morse-message/publisher/kafka"
	"github.com/ytake/morse-message/publisher/log"
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
				Name:    "message:publish",
				Aliases: []string{"m:p"},
				Usage:   "to Kafka",
				Action: func(context *cli.Context) error {
					kc, err := kafka.NewProducer(c.Kafka.BootstrapServers, c.Kafka.ActionCreatedSeparateTopic)
					if err != nil {
						l.Error("kafka producer error", zap.Error(err))
						return err
					}
					cmp := &command.MessagePublisher{Client: kc}
					return cmp.Run(context)
				},
			},
		},
	}
	app.Name = "MorseMessage (do not use for your application)"
	err := app.Run(os.Args)
	if err != nil {
		l.ServerFatal("console error", zap.Error(err))
	}
}
