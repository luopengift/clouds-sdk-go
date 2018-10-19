package main

import (
	"github.com/luopengift/clouds-sdk-go/cmd/multi-clouds-httpd/aws"
	"github.com/luopengift/gohttp"
)

func main() {
	aws.Init()
	app := gohttp.Init()
	app.Route("^/api/v1/aws/ec2$", &aws.Ec2{})
	app.Route("^/api/v1/aws/ec2/tags$", &aws.Tags{})
	app.Route("^/api/v1/aws/rds$", &aws.Rds{})
	app.Route("^/api/v1/aws/autoscaling$", &aws.AutoScaling{})
	app.Route("^/api/v1/clouds/(?P<provider>[_-a-zA-Z0-9]*)/(?P<resource>[a-zA-Z0-9]*)/$", &aws.AWS{})
	app.Run(":8888")
}
