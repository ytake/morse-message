package main

import (
	"github.com/urfave/cli/v2"
	"github.com/ytake/morse-message/publisher/command"
	"github.com/ytake/morse-message/publisher/config"
	"github.com/ytake/morse-message/publisher/console"
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
				Name:    "message:single_partition_publish",
				Aliases: []string{"m:spp"},
				Usage:   "to Kafka",
				Action: func(context *cli.Context) error {
					pc, err := pub.NewProducer(c.Kafka.KafkaBootstrapServers())
					if err != nil {
						l.Error("pub producer error", zap.Error(err))
						return err
					}
					cmp := &console.SinglePartitionPublisher{
						Command: command.UserAction{
							Client: pub.NewSinglePartitionClient(c.Kafka.SingleUserActionTopic(), pc)}}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:single_partition_subscriber",
				Aliases: []string{"m:sps"},
				Usage:   "from Kafka",
				Action: func(context *cli.Context) error {
					sc, err := sub.NewConsumer(c.Kafka.KafkaBootstrapServers(), "no_key_subscriber_sample")
					if err != nil {
						l.Error("subscriber error", zap.Error(err))
						return err
					}
					cmp := &console.SinglePartitionSubscriber{
						Client: sub.NewSinglePartitionClient(c.Kafka.SingleUserActionTopic(), sc)}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:no_key_partition_publish",
				Aliases: []string{"m:nkpp"},
				Usage:   "to Kafka",
				Action: func(context *cli.Context) error {
					pc, err := pub.NewProducer(c.Kafka.KafkaBootstrapServers())
					if err != nil {
						l.Error("pub producer error", zap.Error(err))
						return err
					}
					cmp := &console.NoKeyPartitionPublisher{
						Command: command.UserAction{
							Client: pub.NewNoKeyPartitionClient(c.Kafka.NoKeyUserActionTopic(), pc)}}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:no_key_partition_subscriber",
				Aliases: []string{"m:nkps"},
				Usage:   "from Kafka",
				Action: func(context *cli.Context) error {
					sc, err := sub.NewConsumer(c.Kafka.KafkaBootstrapServers(), "no_key_partition_subscriber_sample")
					if err != nil {
						l.Error("subscriber error", zap.Error(err))
						return err
					}
					cmp := &console.MultiplePartitionSubscriber{
						Client: sub.NewPartitionClient(c.Kafka.NoKeyUserActionTopic(), sc)}
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
