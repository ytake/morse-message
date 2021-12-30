package sub

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ytake/morse-message/publisher/message"
	"github.com/ytake/morse-message/publisher/stream"
	"os"
	"os/signal"
	"syscall"
)

// Client Kafka client struct
type Client struct {
	Consumer *kafka.Consumer
}

type Receiver struct {
	subscriber message.Subscriber
}

// NewConsumer make kafka consumer
func NewConsumer(servers, group string) (*Client, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               servers,
		"broker.address.family":           "v4",
		"group.id":                        group,
		"session.timeout.ms":              6000,
		// "enable.auto.commit":              "false", // 頭から実行したい場合にどうぞ
		"auto.offset.reset":               "earliest",
		"go.application.rebalance.enable": true,
		"enable.partition.eof":            true,
		"go.events.channel.enable":        true,
	})
	return &Client{Consumer: c}, err
}

func (c *Receiver) Subscribe(reader stream.Reader) error {
	topics := []string{c.subscriber.RetrieveTopic()}
	err := c.subscriber.Client().SubscribeTopics(topics,
		func(c *kafka.Consumer, event kafka.Event) error {
			fmt.Sprintf("rebalanced: %s", event)
			return nil
		})
	defer c.subscriber.Client().Close()
	if err != nil {
		return err
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-c.subscriber.Client().Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.subscriber.Client().Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.subscriber.Client().Unassign()
			case *kafka.Message:
				fmt.Fprintf(os.Stderr, "partition %s", e.TopicPartition)
				if err := reader.Proceed(e.Value); err != nil {
					// logger
				}
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Println(e.String())
				// logger
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Println(e.String())
				// logger
			}
		}
	}
	return nil
}
