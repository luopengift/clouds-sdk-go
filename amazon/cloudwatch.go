package amazon

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// ListMertics list mertics
func ListMertics(ctx context.Context, sess *session.Session, filters map[string]string) ([]*cloudwatch.Metric, error) {
	svc := cloudwatch.New(sess)
	input := &cloudwatch.ListMetricsInput{
		Namespace:  aws.String("AWS/Kinesis"),
		MetricName: aws.String("PutRecords.Records"),
	}
	var res []*cloudwatch.Metric
	err := svc.ListMetricsPagesWithContext(ctx, input, func(output *cloudwatch.ListMetricsOutput, ok bool) bool {
		for _, out := range output.Metrics {
			res = append(res, out)
		}
		return output.NextToken != aws.String("")
	})
	return res, err
}

// GetMetricStatistics get metrics
func GetMetricStatistics(ctx context.Context, sess *session.Session, filters map[string]string) (*cloudwatch.GetMetricStatisticsOutput, error) {
	res, err := ListKinesis(ctx, sess, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(res)
	var rest []*cloudwatch.Dimension
	for i, r := range res {
		if i == 10 {
			break
		}
		rest = append(rest, &cloudwatch.Dimension{
			Name:  aws.String("StreamName"),
			Value: aws.String(r),
		})

	}
	// res, err := ListMertics(ctx, sess, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// var des []*cloudwatch.Dimension
	// for _, mertic := range res {
	// 	des = append(des, mertic.Dimensions...)
	// }
	// fmt.Println(des, len(des))
	svc := cloudwatch.New(sess)
	input := &cloudwatch.GetMetricStatisticsInput{
		//Dimensions:         rest,
		Namespace:          aws.String("AWS/Kinesis"),
		MetricName:         aws.String("PutRecords.Records"),
		StartTime:          aws.Time(time.Now().Add(-24 * time.Hour)),
		EndTime:            aws.Time(time.Now()),
		Period:             aws.Int64(600),
		ExtendedStatistics: []*string{aws.String("p100")},
	}
	return svc.GetMetricStatistics(input)
}
