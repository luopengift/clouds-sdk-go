package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/framework"
	"github.com/luopengift/gohttp"
)

// AWS AWS
type AWS struct {
	gohttp.APIHandler
	Sessions []*session.Session
	context.Context
}

// Prepare prepare
func (aws *AWS) Prepare() {
	aws.Sessions = sessions
	aws.Context = context.Background()
}

var sessions []*session.Session

// Init init
func Init() {
	for _, region := range []string{"cn-northwest-1", "cn-north-1"} {
		sess := amazon.MustSession(map[string]string{"region": region})
		sessions = append(sessions, sess)
	}
}

func init() {
	framework.HttpdRoute("^/api/v1/clouds/(?P<provider>[_-a-zA-Z0-9]*)/(?P<resource>[a-zA-Z0-9]*)/$", &AWS{})
}
