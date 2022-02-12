package pub

import "github.com/confluentinc/confluent-kafka-go/kafka"

// HasKeyPartitionClient for no key
type HasKeyPartitionClient struct {
	kafka     *Client
	topic     *string
	partition int32
}

// NewHasKeyPartitionClient client for no key example
func NewHasKeyPartitionClient(topic string, c *Client) *Messenger {
	return &Messenger{publisher: &HasKeyPartitionClient{kafka: c, topic: &topic}}
}

func (c HasKeyPartitionClient) Client() *kafka.Producer {
	return c.kafka.Producer
}

func (c HasKeyPartitionClient) RetrieveTopic() *string {
	return c.topic
}

func (c HasKeyPartitionClient) RetrievePartition() int32 {
	return kafka.PartitionAny
}
