package log

import (
	"fmt"
	"go.uber.org/zap"
)

type (
	// Logger Logger Interface
	Logger interface {
		Info(args ...interface{})
		Error(args ...interface{})
		ServerFatal(args ...interface{})
		RuntimeFatalError(args ...interface{})
	}
	// Write ログ構造体
	Write struct {
		provider *zap.Logger
	}
)

// NewLogger construct
func NewLogger() *Write {
	logger, _ := zap.NewProduction()
	return &Write{
		provider: logger,
	}
}

// Info InfoLevel
func (l *Write) Info(args ...interface{}) {
	l.provider.Info(fmt.Sprint(args...))
}

// Error E Level
func (l *Write) Error(args ...interface{}) {
	l.provider.Error(fmt.Sprint(args...))
}

// RuntimeFatalError Runtime Error
func (l *Write) RuntimeFatalError(args ...interface{}) {
	l.provider.Fatal(fmt.Sprint(args...))
}

// ServerFatal Fatal Error
func (l *Write) ServerFatal(args ...interface{}) {
	l.provider.Fatal(fmt.Sprint(args...))
}

func (l *Write) Flush() error {
	return l.provider.Sync()
}
