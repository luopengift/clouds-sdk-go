package aws

import (
	"github.com/luopengift/framework"
		"github.com/luopengift/clouds-sdk-go/amazon"

	"github.com/aws/aws-sdk-go/service/elasticache"
)

// Elasticache Elasticache
type Elasticache struct {
	AWS
}

// GET method
func (ctx *Elasticache) GET() {
	var results []*elasticache.CacheCluster
	for _, sess := range ctx.Sessions {
		var res *elasticache.DescribeCacheClustersOutput
		res, ctx.APIOutput.Err = amazon.DescribeElasticCache(ctx.Context, sess, nil)
		if ctx.APIOutput.Err != nil {
			return
		}
		results = append(results, res.CacheClusters...)
	}
	ctx.Data = results
}

// ElasticacheRI ElasticacheRI
type ElasticacheRI struct {
	AWS
}

// GET method
func (ctx *ElasticacheRI) GET() {
	ctx.Data = "TODO"
	// ctx.Data, ctx.APIOutput.Err = amazon.DescribeReservedCacheNodes(ctx.Context, ctx.Session, nil)
}

func init() {
	framework.HttpdRoute("^/api/v1/aws/elasticache", &Elasticache{})
	framework.HttpdRoute("^/api/v1/aws/elasticache/ri", &ElasticacheRI{})
}
