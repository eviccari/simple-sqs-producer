package handlers

import (
	"context"
	"errors"
	"github.com/eviccari/rest-http-utils/httperrors"
	"github.com/eviccari/rest-http-utils/httputils"
	"github.com/eviccari/simple-sqs-producer/adapters"
	"github.com/eviccari/simple-sqs-producer/dtos"
	"github.com/eviccari/simple-sqs-producer/services"
	"net/http"
	"time"
)

type MessageProducerHandler struct {
	ctx     context.Context
	service services.MessageProducerService
	logger  adapters.LoggerAdapter
}

// NewMessageProducerHandler - Create new MessageProducerHandler instance
func NewMessageProducerHandler(ctx context.Context, service services.MessageProducerService, logger adapters.LoggerAdapter) *MessageProducerHandler {
	return &MessageProducerHandler{
		ctx:     ctx,
		service: service,
		logger:  logger,
	}
}

// HandlePOST - HTTP POST handler to receive message data and redirect to the queue
func (mph *MessageProducerHandler) HandlePOST(w http.ResponseWriter, r *http.Request) {
	queueName := httputils.GetValueFromHeader(r, "x-app-queue-name")
	if queueName == "" {
		httputils.WriteJSONErrorResponse(w, httperrors.NewBadRequestError(errors.New("x-app-queue-name is required")))
		return
	}

	ctx, cancel := context.WithTimeout(mph.ctx, time.Second*10)
	defer cancel()

	protocolID, err := mph.service.Send(ctx, "test", queueName)
	if err != nil {
		httputils.WriteJSONErrorResponse(w, httperrors.NewInternalServerError(err))
		return
	}

	httputils.WriteJSONResponse(w, dtos.NewSuccessResponse(protocolID), http.StatusOK)
	return
}
