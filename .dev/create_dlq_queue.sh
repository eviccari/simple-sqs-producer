aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name test-dlq-queue --attributes file://mock_dlq_queue.json