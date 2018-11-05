package main

import (
	"fmt"

	"github.com/luopengift/clouds-sdk-go/cmd/clouds-proxy-httpd/aws"
	"github.com/luopengift/gohttp"
)

func main() {
	aws.Init()
	app := gohttp.Init()
	m := map[string]gohttp.Handler{
		"ec2":         &aws.Ec2{},
		"ri":          &aws.RI{},
		"tags":        &aws.Tags{},
		"rds":         &aws.Rds{},
		"autoscaling": &aws.AutoScaling{},
		"sqs":         &aws.Sqs{},
		"elasticache": &aws.Elasticache{},
	}
	for k, v := range m {
		uri := fmt.Sprintf("^/api/v1/aws/%s$", k)
		app.Route(uri, v)
	}
	app.Route("^/api/v1/clouds/(?P<provider>[_-a-zA-Z0-9]*)/(?P<resource>[a-zA-Z0-9]*)/$", &aws.AWS{})
	app.Run(":8888")
}
