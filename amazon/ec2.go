package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// CreateTags CreateTags
func CreateTags(ctx context.Context, sess *session.Session, resources []string, tags map[string]string) (*ec2.CreateTagsOutput, error) {
	svc := ec2.New(sess)

	input := &ec2.CreateTagsInput{
		Resources: func(resources []string) []*string {
			var res []*string
			for _, resource := range resources {
				res = append(res, aws.String(resource))
			}
			return res
		}(resources),
		Tags: func(tags map[string]string) []*ec2.Tag {
			var res []*ec2.Tag
			for key, value := range tags {
				res = append(res, &ec2.Tag{
					Key:   aws.String(key),
					Value: aws.String(value),
				})
			}
			return res
		}(tags),
	}
	return svc.CreateTagsWithContext(ctx, input)
}

//DeleteTags DeleteTags
func DeleteTags(ctx context.Context, sess *session.Session, resources []string, tags map[string]string) (*ec2.DeleteTagsOutput, error) {
	svc := ec2.New(sess)
	input := &ec2.DeleteTagsInput{
		Resources: func(resources []string) []*string {
			var res []*string
			for _, resource := range resources {
				res = append(res, aws.String(resource))
			}
			return res
		}(resources),
		Tags: func(tags map[string]string) []*ec2.Tag {
			var res []*ec2.Tag
			for key, value := range tags {
				res = append(res, &ec2.Tag{
					Key:   aws.String(key),
					Value: aws.String(value),
				})
			}
			return res
		}(tags),
	}
	return svc.DeleteTagsWithContext(ctx, input)
}

func ec2Filters(filters map[string]string) []*ec2.Filter {
	var res []*ec2.Filter
	for key, value := range filters {
		filter := ec2.Filter{
			Name:   aws.String(key),
			Values: []*string{aws.String(value)},
		}
		res = append(res, &filter)
	}
	return res
}

// DescribeInstances xx
func DescribeInstances(ctx context.Context, sess *session.Session, filters map[string]string) ([]*ec2.Instance, error) {
	var results []*ec2.Instance
	svc := ec2.New(sess)
	inputParams := &ec2.DescribeInstancesInput{
		DryRun:     aws.Bool(false),
		Filters:    ec2Filters(filters),
		MaxResults: aws.Int64(100),
	}
	fn := func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
		for _, reservation := range page.Reservations {
			results = append(results, reservation.Instances...)
		}
		return bool(page.NextToken != aws.String(""))
	}
	if err := svc.DescribeInstancesPagesWithContext(ctx, inputParams, fn); err != nil {
		return nil, err
	}
	return results, nil
}

func DescribeReservedInstances(ctx context.Context, sess *session.Session, filters map[string]string) (*ec2.DescribeReservedInstancesOutput, error) {
	svc := ec2.New(sess)
	inputParams := ec2.DescribeReservedInstancesInput{
		Filters: ec2Filters(filters),
	}
	return svc.DescribeReservedInstancesWithContext(ctx, &inputParams)
}

// DescribeVolumes xx
func DescribeVolumes(ctx context.Context, sess *session.Session, callback func(*ec2.Volume) error) error {
	svc := ec2.New(sess)
	inputParams := &ec2.DescribeVolumesInput{
		MaxResults: aws.Int64(100),
	}
	fn := func(page *ec2.DescribeVolumesOutput, lastPage bool) bool {
		for _, volume := range page.Volumes {
			callback(volume)
		}
		return bool(page.NextToken != aws.String(""))
	}
	if err := svc.DescribeVolumesPagesWithContext(ctx, inputParams, fn); err != nil {
		return err
	}
	return nil
}
