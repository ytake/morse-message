package pub

import "github.com/confluentinc/confluent-kafka-go/kafka"

// SinglePartitionClient for no key
type SinglePartitionClient struct {
	kafka     *Client
	topic     *string
	partition int32
}

// NewSinglePartitionClient client for no key example
func NewSinglePartitionClient(topic string, c *Client) *Messenger {
	return &Messenger{publisher: &SinglePartitionClient{kafka: c, topic: &topic}}
}

func (c SinglePartitionClient) Client() *kafka.Producer {
	return c.kafka.Producer
}

func (c SinglePartitionClient) RetrieveTopic() *string {
	return c.topic
}

func (c SinglePartitionClient) RetrievePartition() int32 {
	return kafka.PartitionAny
}
