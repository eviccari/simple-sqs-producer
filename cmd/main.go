package main

import (
	"context"
	"fmt"
	"github.com/eviccari/simple-sqs-producer/adapters"
	"github.com/eviccari/simple-sqs-producer/configs"
	"github.com/eviccari/simple-sqs-producer/handlers"
	"github.com/eviccari/simple-sqs-producer/infra"
	"github.com/eviccari/simple-sqs-producer/services"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"strconv"

	_ "github.com/eviccari/simple-sqs-producer/docs"
)

// @title Simple SQS Producer API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	ctx := context.Background()
	logger := adapters.NewBasicLoggerAdapter()
	configs.PrintConfig(&logger)

	sess := infra.GetAWSSession()
	mpa := adapters.NewSQSMessagePublisherAdapter(sess, &logger)
	service := services.NewBasicMessageProducerService(mpa, &logger)
	messageProducerHandler := handlers.NewMessageProducerHandler(ctx, service, &logger)
	healthCheckHandler := handlers.NewHealthCheckHandler(ctx, &logger)
	swaggerURL := fmt.Sprintf("http://localhost:%d/swagger/doc.json", configs.GetHTTPPort())

	r := chi.NewRouter()
	r.Post("/", messageProducerHandler.HandlePOST)
	r.Get("/health", healthCheckHandler.HandleGET)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL), //The url pointing to API definition
	))

	httpPort := configs.GetHTTPPort()
	if err := http.ListenAndServe(":"+strconv.Itoa(httpPort), r); err != nil {
		panic(err)
	}
}
