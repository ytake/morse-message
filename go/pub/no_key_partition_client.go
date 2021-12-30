package pub

import "github.com/confluentinc/confluent-kafka-go/kafka"

// NoKeyPartitionClient for no key
type NoKeyPartitionClient struct {
	kafka     *Client
	topic     *string
	partition int32
}

// NewNoKeyPartitionClient client for no key example
func NewNoKeyPartitionClient(topic string, c *Client) *Messenger {
	return &Messenger{publisher: &NoKeyPartitionClient{kafka: c, topic: &topic}}
}

func (c NoKeyPartitionClient) Client() *kafka.Producer {
	return c.kafka.Producer
}

func (c NoKeyPartitionClient) RetrieveTopic() *string {
	return c.topic
}

func (c NoKeyPartitionClient) RetrievePartition() int32 {
	return kafka.PartitionAny
}
