package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/gohttp"
)

type AWS struct {
	gohttp.APIHandler
	*session.Session
	context.Context
}

func (aws *AWS) Prepare() {
	aws.Session = sess
	aws.Context = context.Background()
}

var sess *session.Session

func Init() {
	sess = amazon.MustSession(map[string]string{
		"region": "cn-northwest-1",
	})
}
