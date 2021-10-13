package message

import "github.com/confluentinc/confluent-kafka-go/kafka"

// Publisher interface
type Publisher interface {
	Client() *kafka.Producer
	RetrieveTopic() *string
}
