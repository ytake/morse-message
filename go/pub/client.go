package pub

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ytake/morse-message/publisher/message"
)

// Client Kafka client struct
type Client struct {
	Producer *kafka.Producer
}

// NoKeyClient for no key
type NoKeyClient struct {
	kafka *Client
	topic  *string
}

type Messenger struct {
	publisher message.Publisher
}

// NewProducer create producer
func NewProducer(broker string) (*Client, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"message.timeout.ms": "300000",
		"socket.timeout.ms": "30000",
		"message.send.max.retries": "5",
	})
	return &Client{Producer: p}, err
}

// NewNoKeyClient client for no key example
func NewNoKeyClient(topic string, c *Client) *Messenger {
	return &Messenger{publisher: &NoKeyClient{kafka: c, topic: &topic}}
}

func (c *Messenger) Publish(byte []byte) error {
	deliveryChan := make(chan kafka.Event)
	err := c.publisher.Client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: c.publisher.RetrieveTopic(), Partition: kafka.PartitionAny},
		Value:          byte,
	}, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("failed to deliver message: %v\n",
			m.TopicPartition)
	} else {
		fmt.Printf("delivered to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic,
			m.TopicPartition.Partition,
			m.TopicPartition.Offset)
	}
	return nil
}

// Close a Producer instance.
func (c *Messenger) Close() {
	c.publisher.Client().Close()
}

func (c NoKeyClient) Client() *kafka.Producer {
	return c.kafka.Producer
}

func (c NoKeyClient) RetrieveTopic() *string {
	return c.topic
}
