package aws

import (
	"github.com/luopengift/framework"
)

// Sqs sqs
type Sqs struct {
	AWS
}

// GET method
func (ctx *Sqs) GET() {
	ctx.Data = "TODO"
	// ctx.Data, ctx.APIOutput.Err = amazon.ListSQS(ctx.Context, ctx.Session)
}

// POST method
func (ctx *Sqs) POST() {
	ctx.Data = "TODO"

	// name := ctx.GetQuery("name", "")
	// if name == "" {
	// 	ctx.Set(101, "name is null")
	// 	return
	// }
	// ctx.Data, ctx.APIOutput.Err = amazon.CreateSQS(ctx.Context, ctx.Session, name)
}

func init() {
	framework.HttpdRoute("^/api/v1/aws/sqs", &Sqs{})
}
