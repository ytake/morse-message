package config

import "os"

type App struct {
	Kafka struct {
		BootstrapServers           string
		ActionCreatedSeparateTopic string
		ActionDeletedSeparateTopic string
		ActionCreatedTopic         string
	}
}

// New app config
func New() *App {
	return &App{
		Kafka: struct {
			BootstrapServers           string
			ActionCreatedSeparateTopic string
			ActionDeletedSeparateTopic string
			ActionCreatedTopic         string
		}{
			BootstrapServers:           kafkaBootstrapServers(),
			ActionCreatedSeparateTopic: userActionCreatedSeparateTopic(),
			ActionDeletedSeparateTopic: userActionDeletedSeparateTopic(),
			ActionCreatedTopic:         userActionTopic()},
	}
}

func kafkaBootstrapServers() string {
	return os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
}

func userActionCreatedSeparateTopic() string {
	return "user-action-created"
}

func userActionDeletedSeparateTopic() string {
	return "user-action-deleted"
}

func userActionTopic() string {
	return "created-user-action"
}
