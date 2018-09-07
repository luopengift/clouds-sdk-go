package amazon

import "github.com/aws/aws-sdk-go/service/ec2"

func parseTagsWithEc2(instance *ec2.Instance) map[string]string {
	tags := make(map[string]string)
	for _, tag := range instance.Tags {
		tags[*tag.Key] = *tag.Value
	}
	return tags
}
