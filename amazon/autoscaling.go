package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

// DescribeAutoScalingGroups xx
func DescribeAutoScalingGroups(ctx context.Context, sess *session.Session, filters map[string]string) ([]*autoscaling.Group, error) {
	var results []*autoscaling.Group
	svc := autoscaling.New(sess)
	inputParams := &autoscaling.DescribeAutoScalingGroupsInput{
		MaxRecords: aws.Int64(100),
	}
	fn := func(page *autoscaling.DescribeAutoScalingGroupsOutput, lastPage bool) bool {
		results = append(results, page.AutoScalingGroups...)
		return bool(page.NextToken != aws.String(""))
	}
	if err := svc.DescribeAutoScalingGroupsPagesWithContext(ctx, inputParams, fn); err != nil {
		return nil, err
	}
	return results, nil
}
