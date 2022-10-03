package adapters

import "context"

// MessagePublisherAdapter - Describe MessagePublisherAdapter interface that must be implemented by message publisher services
type MessagePublisherAdapter interface {
	Send(ctx context.Context, message string, queueName string) (protocolID string, err error)
}
