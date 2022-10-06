package services

import (
	"context"
	"fmt"
	"github.com/eviccari/simple-sqs-producer/adapters"
)

// BasicMessageProducerService - Describe BasicMessageProducerService struct
type BasicMessageProducerService struct {
	mpa    adapters.MessagePublisherAdapter
	logger adapters.LoggerAdapter
}

// NewBasicMessageProducerService - Create BasicMessageProducerService instance
func NewBasicMessageProducerService(mpa adapters.MessagePublisherAdapter, logger adapters.LoggerAdapter) *BasicMessageProducerService {
	return &BasicMessageProducerService{
		mpa:    mpa,
		logger: logger,
	}
}

// Send - Send message to queue
func (bmps *BasicMessageProducerService) Send(ctx context.Context, message string, queueName string) (protocolID string, err error) {
	protocolID, err = bmps.mpa.Send(ctx, message, queueName)
	if err != nil {
		bmps.logger.Error(fmt.Sprintf("service layer error: %s", err.Error()))
	}

	return
}
