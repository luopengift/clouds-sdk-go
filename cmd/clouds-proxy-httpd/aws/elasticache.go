package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
)

// Elasticache Elasticache
type Elasticache struct {
	AWS
}

// GET method
func (ctx *Elasticache) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeElasticCache(ctx.Context, ctx.Session, nil)
}

// ElasticacheRI ElasticacheRI
type ElasticacheRI struct {
	AWS
}

// GET method
func (ctx *ElasticacheRI) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeReservedCacheNodes(ctx.Context, ctx.Session, nil)
}
