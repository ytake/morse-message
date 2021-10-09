package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Client Kafka client struct
type Client struct {
	producer *kafka.Producer
	topic    *string
}

// NewProducer create producer
func NewProducer(broker, topic string) (*Client, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"message.timeout.ms": "300000",
		"socket.timeout.ms": "30000",
		"message.send.max.retries": "5",
	})
	return &Client{producer: p, topic: &topic}, err
}

func (c *Client) Publish(byte []byte) error {
	deliveryChan := make(chan kafka.Event)
	err := c.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: c.topic, Partition: kafka.PartitionAny},
		Value:          byte,
	}, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	message := e.(*kafka.Message)
	if message.TopicPartition.Error != nil {
		fmt.Printf("failed to deliver message: %v\n",
			message.TopicPartition)
	} else {
		fmt.Printf("delivered to topic %s [%d] at offset %v\n",
			*message.TopicPartition.Topic,
			message.TopicPartition.Partition,
			message.TopicPartition.Offset)
	}
	return nil
}

// Close a Producer instance.
func (c *Client) Close() {
	c.producer.Close()
}
