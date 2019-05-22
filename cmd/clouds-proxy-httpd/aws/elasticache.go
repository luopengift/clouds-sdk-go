package aws

import (
	"github.com/luopengift/framework"
)

// Elasticache Elasticache
type Elasticache struct {
	AWS
}

// GET method
func (ctx *Elasticache) GET() {
	ctx.Data = "TODO"

	// ctx.Data, ctx.APIOutput.Err = amazon.DescribeElasticCache(ctx.Context, ctx.Session, nil)
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
