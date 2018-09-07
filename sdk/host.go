package sdk

import (
	"context"
	"errors"
)

// Hoster Host er
type Hoster interface {
	DescribeHostTypes() ([]HostTyper, error)                                            // 返回可用实例类型
	RunHosts(context.Context, HostInputer) (HostOutputer, error)                        //创建虚拟机
	DescribeHosts(context.Context, DescribeHostsInputer) (DescribeHostsOutputer, error) // 查询虚拟机

}

type defaultHost struct{}

func (a *defaultHost) DescribeHostTypes() ([]HostTyper, error) {
	return nil, errors.New("Not Support")
}

func (a *defaultHost) RunHosts(context.Context, HostInputer) (HostOutputer, error) {
	return nil, errors.New("Not Support")
}
func (a *defaultHost) DescribeHosts(context.Context, DescribeHostsInputer) (DescribeHostsOutputer, error) {
	return nil, errors.New("Not Support")
}

// Host host
type Host struct {
	Provider string            `json:"provider"`             // 提供商: aws/xxx
	ID       string            `json:"ID"`                   // 唯一标识
	Name     string            `json:"name" valid:"max=128"` // aws: tag_Name字段
	Region   Regioner          `json:"region"`               // 区域，表示资源所在的地域，每个地域包含一个或多个可用区。
	Zone     string            `json:"zone"`                 // 可用区
	Tags     map[string]string `json:"tags"`
}

// DescribeHostsInputer DescribeHostsInputer
type DescribeHostsInputer interface{}

// DescribeHostsInput implement DescribeHostsInputer interface
type DescribeHostsInput struct{}

// DescribeHostsOutputer DescribeHostsOutputer
type DescribeHostsOutputer interface {
	DescribeHosts() []*Host
}

// DescribeHostsOutput implement DescribeHostsOutputer interface
type DescribeHostsOutput struct {
	Hosts []*Host
}

// DescribeHosts implement DescribeHostsOutputer interface
func (s *DescribeHostsOutput) DescribeHosts() []*Host {
	return s.Hosts
}

// HostInputer interface
type HostInputer interface {
	ID() string
}

// HostInput implement HostInputer interface
type HostInput struct {
	ID   string
	Name string
	IP   string
}

// HostOutputer interface
type HostOutputer interface {
	ID() string
}

// HostOutput implement HostOutputer interface
type HostOutput struct {
	ID   string
	Name string
	IP   string
}
