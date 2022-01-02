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

type Messenger struct {
	publisher message.Publisher
}

// RequestParameter for publisher
type RequestParameter struct {
	Byte []byte
	Key  []byte
}

// NewProducer create producer
func NewProducer(broker string) (*Client, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        broker,
		"api.version.request":      "false",
		"message.timeout.ms":       "300000",
		"socket.timeout.ms":        "30000",
		"message.send.max.retries": "5",
	})
	return &Client{Producer: p}, err
}

// Publish to kafka bootstrap server
func (c *Messenger) Publish(parameter RequestParameter) error {
	deliveryChan := make(chan kafka.Event)
	km := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     c.publisher.RetrieveTopic(),
			Partition: c.publisher.RetrievePartition(),
		},
		Value: parameter.Byte,
	}
	if len(parameter.Key) != 0 {
		km.Key = parameter.Key
	}
	err := c.publisher.Client().Produce(km, deliveryChan)
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
