package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
)

type Rds struct {
	AWS
}

func (ctx *Rds) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeDBInstances(ctx.Context, ctx.Session, nil)
}
