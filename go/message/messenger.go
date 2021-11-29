package message

import "github.com/confluentinc/confluent-kafka-go/kafka"

// Publisher interface
type Publisher interface {
	Client() *kafka.Producer
	RetrieveTopic() *string
	RetrievePartition() int32
}

// Subscriber interface
type Subscriber interface {
	Client() *kafka.Consumer
	RetrieveTopic() string
}
