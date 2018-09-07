package aliyun

import (
	"github.com/luopengift/clouds-sdk-go/sdk"
)

// ALiyun aliyun
type ALiyun struct {
	Name            string
	Region          sdk.Regioner
	accessKeyID     string
	secretAccessKey string
	*ECS            // implements Hoster interface
}

// Init aliyun init
func Init(region, accessKeyID, secretAccessKey string) sdk.Manufacturer {
	//ecsClient, err := ecs.NewClientWithAccessKey(region, accessKeyID, secretAccessKey)
	//if err != nil {
	// 异常处理
	//	panic(err)
	//}
	return &ALiyun{}
}

// DescribeRegions implement Manufacturer interface
func (a *ALiyun) DescribeRegions() ([]sdk.Regioner, error) {
	//for _, region := range []string{""} {
	//sdk.Region(region)
	//}
	return nil, sdk.ErrNotImplement
}

// DescribeHostTypes describe  instacnetype
func (a *ALiyun) DescribeHostTypes() ([]sdk.HostTyper, error) {
	return nil, sdk.ErrNotImplement
}
