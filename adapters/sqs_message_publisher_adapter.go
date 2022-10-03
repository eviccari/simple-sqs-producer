package adapters

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSMessagePublisherAdapter struct {
	sess   *session.Session
	logger LoggerAdapter
}

// NewSQSMessagePublisherAdapter - Create new SQSMessagePublisherAdapter instance
func NewSQSMessagePublisherAdapter(sess *session.Session, logger LoggerAdapter) *SQSMessagePublisherAdapter {
	return &SQSMessagePublisherAdapter{
		sess:   sess,
		logger: logger,
	}
}

func (sqsmpa *SQSMessagePublisherAdapter) Send(ctx context.Context, message string, queueName string) (protocolID string, err error) {
	svc := sqs.New(sqsmpa.sess)
	output, err := svc.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": &sqs.MessageAttributeValue{
				DataType:    aws.String("Number"),
				StringValue: aws.String("6"),
			},
		},
		MessageBody: &message,
		QueueUrl:    &queueName,
	})

	if err != nil {
		return "", err
	}

	sqsmpa.logger.Info(output)
	protocolID = *output.MessageId
	return
}
