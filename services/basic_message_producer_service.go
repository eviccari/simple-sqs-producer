package services

import (
	"context"
	"github.com/eviccari/simple-sqs-producer/adapters"
)

// BasicMessageProducerService - Describe BasicMessageProducerService struct
type BasicMessageProducerService struct {
	mpa adapters.MessagePublisherAdapter
}

// NewBasicMessageProducerService - Create BasicMessageProducerService instance
func NewBasicMessageProducerService(mpa adapters.MessagePublisherAdapter) *BasicMessageProducerService {
	return &BasicMessageProducerService{
		mpa: mpa,
	}
}

// Send - Send message to queue
func (bmps *BasicMessageProducerService) Send(ctx context.Context, message string, queueName string) (protocolID string, err error) {
	return bmps.mpa.Send(ctx, message, queueName)
}
