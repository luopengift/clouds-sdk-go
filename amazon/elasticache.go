package amazon

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func DescribeElasticCache(ctx context.Context, sess *session.Session, filters map[string]string) (*elasticache.DescribeCacheClustersOutput, error) {
	svc := elasticache.New(sess)
	input := &elasticache.DescribeCacheClustersInput{}
	return svc.DescribeCacheClustersWithContext(ctx, input)
}
