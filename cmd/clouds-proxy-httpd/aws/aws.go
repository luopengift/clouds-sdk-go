package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/gohttp"
)

// AWS AWS
type AWS struct {
	gohttp.APIHandler
	*session.Session
	context.Context
}

// Prepare prepare
func (aws *AWS) Prepare() {
	aws.Session = sess
	aws.Context = context.Background()
}

var sess *session.Session

// Init init
func Init() {
	sess = amazon.MustSession(map[string]string{
		"region": "cn-northwest-1",
	})
}
