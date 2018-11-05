package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
)

type Elasticache struct {
	AWS
}

func (ctx *Elasticache) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeElasticCache(ctx.Context, ctx.Session, nil)
}
