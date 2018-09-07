package amazon

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/luopengift/log"
)

var (
	accessKeyID     = ""
	secretAccessKey = ""
)

// Session session
type Session struct {
	Regions         []string
	AccessKeyID     string
	SecrteAccessKey string
}

// CreateSessionWithRegion xx
func CreateSessionWithRegion(region, accessKeyID, secretAccessKey string) *session.Session {
	sess := session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	return sess
}

// CreateSession create session
// session Follow the sequence below
// 1. Ec2Role
// 2. AssumeRole
// 3. Env
//     - AWS_REGION=<string>
//     - AWS_ACCESS_KEY_ID=<string>
//     - AWS_SECRET_ACCESS_KEY=<string>
// 4. ~/.aws/credentials
//     - AWS_SDK_LOAD_CONFIG=1
// 5. config id and key
//   config map[string]string{
//    "region": "cn-northwest-1",
//    "role": "arn:aws-cn:iam::1234567890:user/username",
//    "filename": "/Users/username/.aws/credentials",
//    "profile":  "xxxx",
//    "accessKeyID": "xxxx",
//    "secretAccessKey": "xxxx"
//	}
func CreateSession(config map[string]string) (*session.Session, error) {
	var (
		region          = config["region"]
		role            = config["role"]
		filename        = config["filename"]
		profile         = config["profile"]
		accessKeyID     = config["accessKeyID"]
		secretAccessKey = config["secretAccessKey"]
	)
	opts := session.Options{}
	sess := session.Must(session.NewSessionWithOptions(opts))
	if region != "" {
		sess.Config.Region = aws.String(region)
	}
	sess.Config.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	// sess.Handlers.Send.PushFront(func(r *request.Request) {
	// 	log.Info("Request: %s, Operation: %v, Payload: %v",r.ClientInfo.ServiceName, r.Operation, r.Params)
	// })

	sess.Config.Credentials = credentials.NewChainCredentials([]credentials.Provider{
		&ec2rolecreds.EC2RoleProvider{
			// Pass in a custom timeout to be used when requesting
			// IAM EC2 Role credentials.
			Client: ec2metadata.New(sess),

			// Do not use early expiry of credentials. If a non zero value is
			// specified the credentials will be expired early
			ExpiryWindow: 0,
		},
		&stscreds.AssumeRoleProvider{
			Client:   sts.New(sess),
			RoleARN:  role,
			Duration: stscreds.DefaultDuration, // 15 minutes
		},
		&credentials.EnvProvider{}, //credentials.NewEnvCredentials(),
		&credentials.SharedCredentialsProvider{
			Filename: filename, //filename,
			Profile:  profile,  // profile,
		}, //credentials.NewSharedCredentials(filename, profile),
		&credentials.StaticProvider{Value: credentials.Value{
			AccessKeyID:     accessKeyID,
			SecretAccessKey: secretAccessKey,
			SessionToken:    "",
		}}, //credentials.NewStaticCredentials(id, secret, token),
	})
	value, err := sess.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}
	log.Debug("%#v", value)
	return sess, nil
}
