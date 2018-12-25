package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

// DescribeDBInstances xx
func DescribeDBInstances(ctx context.Context, sess *session.Session, filters map[string]string) ([]*rds.DBInstance, error) {
	var results []*rds.DBInstance
	svc := rds.New(sess)
	input := &rds.DescribeDBInstancesInput{
		//Filters:    nil,
		MaxRecords: aws.Int64(20),
	}
	fn := func(page *rds.DescribeDBInstancesOutput, lastPage bool) bool {
		results = append(results, page.DBInstances...)
		return bool(page.Marker != aws.String(""))
	}
	if err := svc.DescribeDBInstancesPagesWithContext(ctx, input, fn); err != nil {
		return nil, err
	}
	return results, nil
}

// DescribeReservedDBInstances describe ri db
func DescribeReservedDBInstances(ctx context.Context, sess *session.Session, filters map[string]string) (*rds.DescribeReservedDBInstancesOutput, error) {
	svc := rds.New(sess)
	input := &rds.DescribeReservedDBInstancesInput{
		// DBInstanceClass:    aws.String("db.t2.micro"),
		// Duration:           aws.String("1y"),
		// MultiAZ:            aws.Bool(false),
		// OfferingType:       aws.String("No Upfront"),
		// ProductDescription: aws.String("mysql"),
	}
	return svc.DescribeReservedDBInstances(input)
}
