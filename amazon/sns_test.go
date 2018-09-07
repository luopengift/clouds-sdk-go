package amazon

import (
	"context"
	"fmt"
	"testing"

	"github.com/luopengift/log"
)

func Test_CreateTopic(t *testing.T) {
	sess, err := CreateSession(nil)
	res, err := ListSQS(context.Background(), sess)
	log.Info("%v, %v", res, err)
	fmt.Println("ddd")
}
