package aws

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/framework"
)

// Rds Rds
type Rds struct {
	AWS
}

// GET method
func (ctx *Rds) GET() {
	var results []*rds.DBInstance
	for _, sess := range ctx.Sessions {
		var res []*rds.DBInstance
		res, ctx.APIOutput.Err = amazon.DescribeDBInstances(ctx.Context, sess, nil)

		if ctx.APIOutput.Err != nil {
			return
		}
		results = append(results, res...)
	}
	ctx.Data = results
}

// RdsRI RdsRI
type RdsRI struct {
	AWS
}

// GET method
func (ctx *RdsRI) GET() {
	ctx.Data = "TODO"
	// ctx.Data, ctx.APIOutput.Err = amazon.DescribeReservedDBInstances(ctx.Context, ctx.Session, nil)
}

func init() {
	framework.HttpdRoute("^/api/v1/aws/rds$", &Rds{})
	framework.HttpdRoute("^/api/v1/aws/rds/ri$", &RdsRI{})
}
