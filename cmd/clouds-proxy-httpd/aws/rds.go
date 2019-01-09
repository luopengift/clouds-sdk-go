package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/framework"
)

// Rds Rds
type Rds struct {
	AWS
}

// GET method
func (ctx *Rds) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeDBInstances(ctx.Context, ctx.Session, nil)
}

// RdsRI RdsRI
type RdsRI struct {
	AWS
}

// GET method
func (ctx *RdsRI) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeReservedDBInstances(ctx.Context, ctx.Session, nil)
}

func init() {
	framework.HttpdRoute("^/api/v1/aws/rds$", &Rds{})
	framework.HttpdRoute("^/api/v1/aws/rds/ri$", &RdsRI{})
}
