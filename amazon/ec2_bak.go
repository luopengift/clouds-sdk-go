package amazon

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/luopengift/clouds-sdk-go/sdk"
	"github.com/luopengift/types"
	"github.com/luopengift/types/validate"
)

// HostFromEC2 convert ec2 to host
func HostFromEC2(instance *ec2.Instance) (*sdk.Host, error) {
	tags := parseTagsWithEc2(instance)
	host := &sdk.Host{
		Provider: "aws",
		ID:       *instance.InstanceId,
		Name:     tags["Name"],
		Region:   sdk.Region(""),
		Zone:     *instance.Placement.AvailabilityZone,
		Tags:     tags,
	}
	if err := validate.Validate(host); err != nil {
		return nil, err
	}
	return host, nil
}

// EC2 instance, implements Hoster interface
type EC2 struct {
	sessions []*session.Session
}

// DescribeHostTypes describe ec2 instacnetype
func (aws EC2) DescribeHostTypes() ([]sdk.HostTyper, error) {
	var _types []sdk.HostTyper
	for _, _type := range strings.Split(instanceType, "\n") {
		t := strings.Split(_type, "|")
		if len(t) != 3 {
			continue
		}
		cpu, err := types.ToInt(t[1])
		if err != nil {
			return nil, err
		}
		mem, err := types.ToFloat64(t[2])
		if err != nil {
			return nil, err
		}
		_types = append(_types, sdk.NewHostType(t[0], cpu, mem))

	}
	return _types, nil
}

// DescribeHosts describe ec2 instance
func (aws EC2) DescribeHosts(ctx context.Context, describeHostsInput sdk.DescribeHostsInputer) (sdk.DescribeHostsOutputer, error) {
	var output = &sdk.DescribeHostsOutput{
		Hosts: []*sdk.Host{},
	}
	// TODO: error is not handle!
	fun := func(instance *ec2.Instance) error {
		host, err := HostFromEC2(instance)
		if err != nil {
			return err
		}
		output.Hosts = append(output.Hosts, host)
		return nil
	}
	for _, session := range aws.sessions {
		DescribeHosts(ctx, session, nil, fun)
	}
	return output, nil
}

// RunHosts Launches the specified Instances.
func (aws EC2) RunHosts(context.Context, sdk.HostInputer) (sdk.HostOutputer, error) {
	return nil, nil
}

// DescribeDisks xx
func (aws EC2) DescribeDisks(context.Context, sdk.DiskInputer) (sdk.DiskOutputer, error) {
	return nil, nil
}

var _ sdk.Hoster = &EC2{}

// DescribeHosts xx
func DescribeHosts(ctx context.Context, sess *session.Session, filters map[string]string, callback func(*ec2.Instance) error) error {
	svc := ec2.New(sess)
	inputParams := &ec2.DescribeInstancesInput{
		DryRun:     aws.Bool(false),
		Filters:    ec2Filters(filters),
		MaxResults: aws.Int64(100),
	}
	fun := func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				callback(instance)
			}
		}
		return bool(page.NextToken != aws.String(""))
	}
	if err := svc.DescribeInstancesPagesWithContext(ctx, inputParams, fun); err != nil {
		return err
	}
	return nil
}
