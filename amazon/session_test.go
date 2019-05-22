package amazon

import (
	"context"
	"testing"

	"github.com/luopengift/log"
)

func Test_Session(t *testing.T) {
	sess, err := CreateSession(nil)
	if err != nil {
		t.Error("err")
	}
	t.Log(sess, err)
}

func Test_DescribeAutoScalingGroups(t *testing.T) {
	sess, err := CreateSession(map[string]string{
		"region": "cn-northwest-1",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(sess)
	ctx := context.Background()
	res, err := DescribeAutoScalingGroups(ctx, sess, nil)
	if err != nil {
		t.Error(err)
	}
	log.Infof("%v", res)
}
