package main

import (
	"context"
	"github.com/eviccari/simple-sqs-producer/adapters"
	"github.com/eviccari/simple-sqs-producer/configs"
	"github.com/eviccari/simple-sqs-producer/handlers"
	"github.com/eviccari/simple-sqs-producer/infra"
	"github.com/eviccari/simple-sqs-producer/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func main() {

	ctx := context.Background()
	logger := adapters.NewBasicLoggerAdapter()
	configs.PrintConfig(&logger)

	sess := infra.GetAWSSession()
	mpa := adapters.NewSQSMessagePublisherAdapter(sess, &logger)
	service := services.NewBasicMessageProducerService(mpa)
	handler := handlers.NewMessageProducerHandler(ctx, service, &logger)

	r := chi.NewRouter()
	r.Post("/", handler.HandlePOST)

	httpPort := configs.GetHTTPPort()
	http.ListenAndServe(":"+strconv.Itoa(httpPort), r)
}
