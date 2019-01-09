package aws

import (
	"encoding/json"
	"strings"

	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/framework"
)

// Ec2 ec2
type Ec2 struct {
	AWS
}

// GET method
func (ctx *Ec2) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeInstances(ctx.Context, ctx.Session, nil)
}

// Ec2RI Ec2RI
type Ec2RI struct {
	AWS
}

// GET method
func (ctx *Ec2RI) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeReservedInstances(ctx.Context, ctx.Session, nil)
}

// Ec2Tags Ec2Tags
type Ec2Tags struct {
	AWS
}

// POST method
func (ctx *Ec2Tags) POST() {
	resources := ctx.GetQuery("resources", "")
	tags := make(map[string]string)
	ctx.APIOutput.Err = json.Unmarshal(ctx.GetBodyArgs(), &tags)
	if ctx.APIOutput.Err != nil {
		ctx.Set(101, "post body is not json!")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.CreateTags(ctx.Context, ctx.Session, strings.Split(resources, ","), tags)
}

// DELETE method
func (ctx *Ec2Tags) DELETE() {
	resources := ctx.GetQuery("resources", "")
	tags := make(map[string]string)
	ctx.APIOutput.Err = json.Unmarshal(ctx.GetBodyArgs(), &tags)
	if ctx.APIOutput.Err != nil {
		ctx.Set(101, "post body is not json!")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.DeleteTags(ctx.Context, ctx.Session, strings.Split(resources, ","), tags)
}

func init() {
	framework.HttpdRoute("^/api/v1/aws/ec2$", &Ec2{})
	framework.HttpdRoute("^/api/v1/aws/ec2/ri$", &Ec2RI{})
	framework.HttpdRoute("^/api/v1/aws/ec2/tag", &Ec2Tags{})
}
