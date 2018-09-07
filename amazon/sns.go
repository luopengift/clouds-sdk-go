package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// CreateTopic xx
func CreateTopic(ctx context.Context, sess *session.Session, name string) (*sns.CreateTopicOutput, error) {
	svc := sns.New(sess)
	inputParams := sns.CreateTopicInput{
		Name: aws.String(name),
	}
	return svc.CreateTopicWithContext(ctx, &inputParams)
}
