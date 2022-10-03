package services

import "context"

// MessageProducerService - Describe interface that must be implemented by message producer services
type MessageProducerService interface {
	Send(ctx context.Context, message string, queueName string) (protocolID string, err error)
}
