package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/luopengift/clouds-sdk-go/cmd/clouds-proxy-httpd/aws"
	"github.com/luopengift/framework"
	"github.com/luopengift/gohttp"
)

func main() {
	ctx := context.Background()
	app := framework.New()
	app.InitFunc(func(ctx context.Context) error {
		aws.Init()
		m := map[string]gohttp.Handler{
			"ec2":            &aws.Ec2{},
			"ec2/ri":         &aws.Ec2RI{},
			"ec2/tag":        &aws.Ec2Tags{},
			"rds":            &aws.Rds{},
			"rds/ri":         &aws.RdsRI{},
			"autoscaling":    &aws.AutoScaling{},
			"sqs":            &aws.Sqs{},
			"elasticache":    &aws.Elasticache{},
			"elasticache/ri": &aws.ElasticacheRI{},
		}
		for k, v := range m {
			uri := fmt.Sprintf("^%s$", filepath.Join("/api/v1/aws", k))
			app.Route(uri, v)
		}
		app.Route("^/api/v1/clouds/(?P<provider>[_-a-zA-Z0-9]*)/(?P<resource>[a-zA-Z0-9]*)/$", &aws.AWS{})
		return nil
	})
	app.MainFunc(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
		}
		return nil
	})
	app.Run(ctx)
}
