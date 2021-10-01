package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Client Kafka client struct
type Client struct {
	Producer *kafka.Producer
	Topic    *string
}

// NewProducer create producer
func NewProducer(broker string) (*Client, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	return &Client{Producer: p}, err
}
