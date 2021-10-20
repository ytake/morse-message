package config

import "os"

type (
	Kafka struct{}
	App   struct {
		Kafka Kafka
	}
)

// New app config
func New() *App {
	return &App{
		Kafka: Kafka{},
	}
}

func (k Kafka) KafkaBootstrapServers() string {
	return os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}

func (k Kafka) UserActionCreatedSeparateTopic() string {
	return "user-action-created"
}

func (k Kafka) UserActionDeletedSeparateTopic() string {
	return "user-action-deleted"
}

func (k Kafka) UserActionTopic() string {
	return "created-user-action"
}

func (k Kafka) SingleUserActionTopic() string {
	return "single-created-user-action"
}
