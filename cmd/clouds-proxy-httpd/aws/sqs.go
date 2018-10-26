package aws

import (
	"github.com/luopengift/clouds-sdk-go/amazon"
)

type Sqs struct {
	AWS
}

func (ctx *Sqs) GET() {
	ctx.Data, ctx.APIOutput.Err = amazon.ListSQS(ctx.Context, ctx.Session)
}
func (ctx *Sqs) POST() {
	name := ctx.GetQuery("name", "")
	if name == "" {
		ctx.Set(101, "name is null")
		return
	}
	ctx.Data, ctx.APIOutput.Err = amazon.CreateSQS(ctx.Context, ctx.Session, name)
}
