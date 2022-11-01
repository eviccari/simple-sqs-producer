package handlers

import (
	"context"
	"fmt"
	"github.com/eviccari/rest-http-utils/httputils"
	"github.com/eviccari/simple-sqs-producer/adapters"
	"github.com/eviccari/simple-sqs-producer/dtos"
	"net/http"
)

// HealthCheckHandler - Describe HealthCheckHandler struct
type HealthCheckHandler struct {
	ctx    context.Context
	logger adapters.LoggerAdapter
}

// NewHealthCheckHandler - Create new HealthCheckHandler instance
func NewHealthCheckHandler(ctx context.Context, logger adapters.LoggerAdapter) *HealthCheckHandler {
	return &HealthCheckHandler{
		ctx:    ctx,
		logger: logger,
	}
}

// HandleGET - Verify if service is alive
// PathParamsExample godoc
// @Summary      Verify if service is running. Must be used by container orchestrator
// @Method       GET
// @Description  Verify if service is running. Must be used by container orchestrator
// @Tags         Health Check
// @Produce      json
// @Success      200         {string}  string  "message"
// @Router       /health [get]
func (hch *HealthCheckHandler) HandleGET(w http.ResponseWriter, r *http.Request) {
	hch.logger.Info(fmt.Sprintf("origin: %s, check system health: OK", r.Host))
	httputils.WriteJSONResponse(w, dtos.NewHealthCheckDTO(), http.StatusOK)
	return
}
