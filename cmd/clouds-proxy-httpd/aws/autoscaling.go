package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
)

// AutoScaling AutoScaling
type AutoScaling struct {
	AWS
}

// GET method
func (ctx *AutoScaling) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeAutoScalingGroups(ctx.Context, ctx.Session, nil)
}
