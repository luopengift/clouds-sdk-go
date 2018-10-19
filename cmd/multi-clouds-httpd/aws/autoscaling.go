package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
)

type AutoScaling struct {
	AWS
}

func (ctx *AutoScaling) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeAutoScalingGroups(ctx.Context, ctx.Session, nil)
}
