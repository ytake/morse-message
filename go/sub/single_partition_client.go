package sub

import "github.com/confluentinc/confluent-kafka-go/kafka"

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
