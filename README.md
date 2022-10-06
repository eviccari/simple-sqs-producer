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
````bash
swag init -g handlers/*.go   
````

## Run Application
Locally, application needs the Localstack endpoint to produce messages. It can be provided by [run_local_stack.sh](./.dev/run_local_stack.sh) file.
The script will launch docker container with mocked AWS services including SQS.

AWS Credential Configurations is needed to load AWS SDK and establish connection (even with localstack). The basic instructions can be found at [AWS CLI configuration Page](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html).

To create simple Test Queue and Test Dead Letter Queue (DLQ), run [create_dlq_queue.sh](./.dev/create_dlq_queue.sh) and [create_queue.sh](./.dev/create_queue.sh) in that order (int the .dev folder):
````bash
./create_dlq_queue.sh
./create_queue.sh
````

The results look's like bellow:
````json
{
    "QueueUrl": "http://localhost:4566/000000000000/test-dlq-queue"
}
````
````json
{
  "QueueUrl": "http://localhost:4566/000000000000/test-queue"
}
````

The **test-queue** URL could be used 


Compile and execute main process (in the root project folder):
````bash
go run ./cmd/main.go
````




