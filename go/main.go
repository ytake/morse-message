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

const (
	partitionFlagName = "partition"
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
					cmp := &command.SinglePartitionPublisher{
						Client: pub.NewSinglePartitionClient(c.Kafka.SingleUserActionTopic(), pc)}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:single_partition_subscriber",
				Aliases: []string{"m:sps"},
				Usage:   "from Kafka",
				Action: func(context *cli.Context) error {
					sc, err := sub.NewConsumer(c.Kafka.KafkaBootstrapServers(), "nokey_subscriber_sample")
					if err != nil {
						l.Error("subscriber error", zap.Error(err))
						return err
					}
					cmp := &command.SinglePartitionSubscriber{
						Client: sub.NewSinglePartitionClient(c.Kafka.SingleUserActionTopic(), sc)}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:no_key_partition_publish",
				Aliases: []string{"m:nkpp"},
				Usage:   "to Kafka",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     partitionFlagName,
						Usage:    "for kafka partition (0 or 1?)",
						Required: true,
					},
				},
				Action: func(context *cli.Context) error {
					pc, err := pub.NewProducer(c.Kafka.KafkaBootstrapServers())
					if err != nil {
						l.Error("pub producer error", zap.Error(err))
						return err
					}
					p := context.Int(partitionFlagName)
					cmp := &command.NoKeyPartitionPublisher{
						Client: pub.NewNoKeyPartitionClient(
							c.Kafka.NoKeyUserActionTopic(),
							int32(p),
							pc)}
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
