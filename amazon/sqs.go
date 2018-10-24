package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// CreateSQS create sqs
func CreateSQS(ctx context.Context, sess *session.Session, name string) (*sqs.CreateQueueOutput, error) {
	svc := sqs.New(sess)
	inputParams := sqs.CreateQueueInput{
		QueueName: aws.String(name),
	}
	return svc.CreateQueue(&inputParams)
}

//ListSQS list sqs
func ListSQS(ctx context.Context, sess *session.Session) (*sqs.ListQueuesOutput, error) {
	svc := sqs.New(sess)
	return svc.ListQueues(&sqs.ListQueuesInput{})
}

// Send send sqs
func Send(ctx context.Context, sess *session.Session, url, msg string) (*sqs.SendMessageOutput, error) {
	svc := sqs.New(sess)
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(msg),
		QueueUrl:    aws.String(url),
	}
	return svc.SendMessageWithContext(ctx, input)
}

// Receive receive sqs
func Receive(ctx context.Context, sess *session.Session, url string) (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(sess)
	input := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(url),
	}
	return svc.ReceiveMessageWithContext(ctx, input)
}
