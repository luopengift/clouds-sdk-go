package aws

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/framework"
)

// AutoScaling AutoScaling
type AutoScaling struct {
	AWS
}

// GET method
func (ctx *AutoScaling) GET() {
	var results []*autoscaling.Group
	for _, sess := range ctx.Sessions {
		var res []*autoscaling.Group
		res, ctx.APIOutput.Err = amazon.DescribeAutoScalingGroups(ctx.Context, sess, nil)
		if ctx.APIOutput.Err != nil {
			return
		}
		results = append(results, res...)
	}
	ctx.Data = results
}

func init() {
	framework.HttpdRoute("^/api/v1/aws/autoscaling$", &AutoScaling{})
}
