package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eviccari/rest-http-utils/httperrors"
	"github.com/eviccari/rest-http-utils/httputils"
	"github.com/eviccari/simple-sqs-producer/adapters"
	"github.com/eviccari/simple-sqs-producer/dtos"
	"github.com/eviccari/simple-sqs-producer/payloads"
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
// PathParamsExample godoc
// @Summary      Endpoint to send message to SQS queue.
// @Description  Endpoint to send message to SQS queue. The queue URL must be sent into <b>x-app-queue-name</b>
// @Tags         SQS Payload
// @Accept       json
// @Produce      json
// @Param        x-app-queue-url header string true "Queue URL"
// @Success      200         {string}  string  "answer"
// @Failure      400         {string}  string  "ok"
// @Failure      404         {string}  string  "ok"
// @Failure      500         {string}  string  "ok"
// @Router       / [post]
func (mph *MessageProducerHandler) HandlePOST(w http.ResponseWriter, r *http.Request) {
	queueName := httputils.GetValueFromHeader(r, "x-app-queue-url")
	if queueName == "" {
		httputils.WriteJSONErrorResponse(w, httperrors.NewBadRequestError(errors.New("x-app-queue-name is required")))
		return
	}

	var inputMessagePayload payloads.InputMessagePayload
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&inputMessagePayload); err != nil {
		httputils.WriteJSONErrorResponse(w, httperrors.NewBadRequestError(err))
		return
	}

	ctx, cancel := context.WithTimeout(mph.ctx, time.Second*10)
	defer cancel()

	protocolID, err := mph.service.Send(ctx, inputMessagePayload.Message, queueName)
	if err != nil {
		httputils.WriteJSONErrorResponse(w, httperrors.NewInternalServerError(err))
		return
	}

	httputils.WriteJSONResponse(w, dtos.NewSuccessResponse(protocolID), http.StatusOK)
	return
}
