package sub

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// SinglePartitionClient for no key
type SinglePartitionClient struct {
	kafka *Client
	topic string
}

// NewSinglePartitionClient client for no key example
func NewSinglePartitionClient(topic string, c *Client) *Receiver {
	return &Receiver{subscriber: &SinglePartitionClient{kafka: c, topic: topic}}
}

func (c SinglePartitionClient) Client() *kafka.Consumer {
	return c.kafka.Consumer
}

func (c SinglePartitionClient) RetrieveTopic() string {
	return c.topic
}

// PartitionClient for no key
type PartitionClient struct {
	kafka *Client
	topic string
}

// NewPartitionClient client for no key example
func NewPartitionClient(topic string, c *Client) *Receiver {
	return &Receiver{subscriber: &PartitionClient{kafka: c, topic: topic}}
}

func (c PartitionClient) Client() *kafka.Consumer {
	return c.kafka.Consumer
}

func (c PartitionClient) RetrieveTopic() string {
	return c.topic
}
