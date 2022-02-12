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
				Usage:   "send a message to topic single-user-action",
				Description: "1partition 1topic構成でproduceする例",
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
				Usage:   "consume topic single-user-action",
				Description: "1partition 1topic構成でconsumeする例",
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
				Usage:   "send a message to topic nokey-user-action",
				Description: "2partition 1topic構成でproduceする例 / keyなしでランダムに分割されるのでうまく格納されない例",
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
				Usage:   "consume topic nokey-user-action",
				Description: "2partition 1topic構成でconsumeする例 / keyなしでランダムに分割されるのでうまく取得できない例",
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
			{
				Name:    "message:has_key_partition_publish",
				Aliases: []string{"m:hkpp"},
				Usage:   "send a message to topic haskey-user-action",
				Description: "メッセージの時間軸をuserごとに担保することで想定通りに格納するサンプル",
				Action: func(context *cli.Context) error {
					pc, err := pub.NewProducer(c.Kafka.KafkaBootstrapServers())
					if err != nil {
						l.Error("pub producer error", zap.Error(err))
						return err
					}
					cmp := &console.HasKeyPartitionPublisher{
						Command: command.UserAction{
							Client: pub.NewHasKeyPartitionClient(c.Kafka.HasKeyUserActionTopic(), pc)}}
					return cmp.Run(context)
				},
			},
			{
				Name:    "message:has_key_partition_subscriber",
				Aliases: []string{"m:hkps"},
				Usage:   "consume topic haskey-user-action",
				Description: "メッセージの時間軸をuserごとに格納したものを想定通りに取得するサンプル",
				Action: func(context *cli.Context) error {
					sc, err := sub.NewConsumer(c.Kafka.KafkaBootstrapServers(), "has_key_partition_subscriber_sample")
					if err != nil {
						l.Error("subscriber error", zap.Error(err))
						return err
					}
					cmp := &console.MultiplePartitionSubscriber{
						Client: sub.NewPartitionClient(c.Kafka.HasKeyUserActionTopic(), sc)}
					return cmp.Run(context)
				},
			},
		},
	}
	app.Name = "SplitMessage (do not use for your application)"
	err := app.Run(os.Args)
	if err != nil {
		l.ServerFatal("console error", zap.Error(err))
	}
}
