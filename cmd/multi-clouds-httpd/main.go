package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/gohttp"
)

var sessMap = map[string]*session.Session{
	"aws": amazon.MustSession(map[string]string{
		"region": "cn-northwest-1",
	}),
	"aliyun": nil,
}

type AWS struct {
	gohttp.APIHandler
	*session.Session
	context.Context
}

func (aws *AWS) Prepare() {
	aws.Session = sessMap["aws"]
	aws.Context = context.Background()
}

type ec2 struct {
	AWS
}

func (ctx *ec2) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.DescribeInstances(ctx.Context, ctx.Session, nil)
}

type tags struct {
	AWS
}

func (ctx *tags) POST() {
	resources := ctx.GetQuery("resources", "")
	tags := make(map[string]string)
	ctx.APIOutput.Err = json.Unmarshal(ctx.GetBodyArgs(), &tags)
	if ctx.APIOutput.Err != nil {
		ctx.Set(101, "post body is not json!")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.CreateTags(ctx.Context, ctx.Session, strings.Split(resources, ","), tags)
}

func (ctx *tags) DELETE() {
	resources := ctx.GetQuery("resources", "")
	tags := make(map[string]string)
	ctx.APIOutput.Err = json.Unmarshal(ctx.GetBodyArgs(), &tags)
	if ctx.APIOutput.Err != nil {
		ctx.Set(101, "post body is not json!")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.DeleteTags(ctx.Context, ctx.Session, strings.Split(resources, ","), tags)
}

func main() {
	app := gohttp.Init()
	app.Route("^/api/v1/aws/ec2$", &ec2{})
	app.Route("^/api/v1/aws/ec2/tags$", &tags{})
	app.Route("^/api/v1/clouds/(?P<provider>[_-a-zA-Z0-9]*)/(?P<resource>[a-zA-Z0-9]*)/$", &AWS{})
	app.Run()
}
