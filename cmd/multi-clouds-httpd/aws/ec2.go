package aws

import (
	"encoding/json"
	"strings"

	"github.com/luopengift/clouds-sdk-go/amazon"
)

type Ec2 struct {
	AWS
}

func (ctx *Ec2) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeInstances(ctx.Context, ctx.Session, nil)
}

type Tags struct {
	AWS
}

func (ctx *Tags) POST() {
	resources := ctx.GetQuery("resources", "")
	tags := make(map[string]string)
	ctx.APIOutput.Err = json.Unmarshal(ctx.GetBodyArgs(), &tags)
	if ctx.APIOutput.Err != nil {
		ctx.Set(101, "post body is not json!")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.CreateTags(ctx.Context, ctx.Session, strings.Split(resources, ","), tags)
}

func (ctx *Tags) DELETE() {
	resources := ctx.GetQuery("resources", "")
	tags := make(map[string]string)
	ctx.APIOutput.Err = json.Unmarshal(ctx.GetBodyArgs(), &tags)
	if ctx.APIOutput.Err != nil {
		ctx.Set(101, "post body is not json!")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.DeleteTags(ctx.Context, ctx.Session, strings.Split(resources, ","), tags)
}
