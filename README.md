# Simple SQS Producer

REST API that sends string message to provided SQS queue

## Technologies
- [AWS SDK](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-sns-with-go-sdk.html)
- [Go-Chi Microframework v5](https://github.com/go-chi/chi)
- [Golang 1.19](https://go.dev/)
- [Localstack](https://docs.localstack.cloud/get-started/)

## Generate Swagger Docs

Install swag cli:
````bash
go install github.com/swaggo/swag/cmd/swag@latest
````

Add swag to project dependencies:
````bash
go get github.com/swaggo/swag/cmd/swag
````

Generate swagger artifacts (in the root project folder):
`````bash
swag init -g handlers/*.go   
`````



