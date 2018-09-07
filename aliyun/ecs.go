package aliyun

import (
	"context"

	"github.com/luopengift/clouds-sdk-go/sdk"
)

// ECS aliyun host, implements sdk.Hoster interface
type ECS struct {
}

// DescribeHostTypes implement Hoster interface
func (e *ECS) DescribeHostTypes() ([]sdk.HostTyper, error) {
	return nil, sdk.ErrNotImplement
}

// RunHosts implement Hoster interface
func (e *ECS) RunHosts(context.Context, sdk.HostInputer) (sdk.HostOutputer, error) {
	return nil, sdk.ErrNotImplement
}

// DescribeHosts implement Hoster interface
func (e *ECS) DescribeHosts(context.Context, sdk.DescribeHostsInputer) (sdk.DescribeHostsOutputer, error) {
	return nil, sdk.ErrNotImplement
}
