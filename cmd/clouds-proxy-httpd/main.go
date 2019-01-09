package main

import (
	"github.com/luopengift/clouds-sdk-go/cmd/clouds-proxy-httpd/aws"
	"github.com/luopengift/framework"
)

func main() {
	aws.Init()
	framework.Run()
}
