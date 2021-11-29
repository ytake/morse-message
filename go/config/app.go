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

// KafkaBootstrapServers for kafka brokers
func (k Kafka) KafkaBootstrapServers() string {
	return os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}

// UserActionCreatedSeparateTopic アクションごとに分割したtopic例
func (k Kafka) UserActionCreatedSeparateTopic() string {
	return "user-action-created"
}

// UserActionDeletedSeparateTopic アクションごとに分割したtopic例
func (k Kafka) UserActionDeletedSeparateTopic() string {
	return "user-action-deleted"
}

// NoKeyUserActionTopic keyを指定せずにpartitionを2以上にしたtopic例
func (k Kafka) NoKeyUserActionTopic() string {
	return "nokey-user-action"
}

// HasKeyUserActionTopic keyを指定しpartitionを2以上にしたtopic例
func (k Kafka) HasKeyUserActionTopic() string {
	return "haskey-user-action"
}

// SingleUserActionTopic partitionを利用せずに1topicで処理するtopic例
func (k Kafka) SingleUserActionTopic() string {
	return "single-user-action"
}
