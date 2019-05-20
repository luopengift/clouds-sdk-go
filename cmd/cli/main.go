package main

import (
	"context"

	"github.com/luopengift/clouds-sdk-go/amazon"
	"github.com/luopengift/log"
)

func main() {
	ctx := context.Background()
	sess := amazon.MustSession(map[string]string{
		"region": "cn-northwest-1",
	})
	res, err := amazon.GetMetricStatistics(ctx, sess, nil)
	log.Info("%#v, %v", res, err)
}
