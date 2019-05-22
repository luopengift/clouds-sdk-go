package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// ListKinesis list kinesis
func ListKinesis(ctx context.Context, sess *session.Session, filters map[string]string) ([]string, error) {
	var res []string
	svc := kinesis.New(sess)
	input := &kinesis.ListStreamsInput{
		Limit: aws.Int64(100),
	}
	err := svc.ListStreamsPagesWithContext(ctx, input, func(output *kinesis.ListStreamsOutput, ok bool) bool {
		for _, name := range output.StreamNames {
			res = append(res, aws.StringValue(name))
		}
		return aws.BoolValue(output.HasMoreStreams)
	})
	return res, err
}
