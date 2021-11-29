package sub

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ytake/morse-message/publisher/message"
	"github.com/ytake/morse-message/publisher/stream"
	"log"
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
		"session.timeout.ms":              30000,
		"enable.auto.commit":              "false", // 頭から実行したい場合にどうぞ
		"auto.offset.reset":               "earliest",
		"go.application.rebalance.enable": true,
		"enable.partition.eof":            true,
		"go.events.channel.enable":        true,
	})
	return &Client{Consumer: c}, err
}

func (c *Receiver) Subscribe(reader stream.Reader) error {
	err := c.subscriber.Client().Subscribe(c.subscriber.RetrieveTopic(),
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
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.subscriber.Client().Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				if e.Headers != nil {
					// logger
				}
				if err := reader.Proceed(e.Value); err != nil {
					// logger
				}
			case kafka.Error:
				// logger
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				// logger
			}
		}
	}
	return nil
}
