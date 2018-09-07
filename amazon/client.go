package amazon

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/luopengift/clouds-sdk-go/sdk"
)

// AWS aws
type AWS struct {
	Name            string
	Region          sdk.Regioner
	accessKeyID     string
	secretAccessKey string
	*EC2            // implements Hoster interface
}

// Init init
func Init(region, accessKeyID, secretAccessKey string) sdk.Manufacturer {
	sess := CreateSessionWithRegion(region, accessKeyID, secretAccessKey)
	return &AWS{
		Name:            "aws",
		Region:          sdk.Region(region),
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		EC2: &EC2{
			sessions: []*session.Session{sess},
		},
	}
}

func (a *AWS) String() string {
	return a.Name + "/" + a.Region.String()
}

// DescribeRegions implement Manufacturer interface
func (a *AWS) DescribeRegions() ([]sdk.Regioner, error) {
	var regions []sdk.Regioner
	for _, region := range []string{"cn-north-1"} {
		regions = append(regions, sdk.Region(region))
	}
	return regions, nil
}
