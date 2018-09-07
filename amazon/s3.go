package amazon

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Upload to s3, origin
func Upload(ctx context.Context, sess *session.Session, src, dest string) (string, error) {
	f, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", src, err)
	}

	u, err := url.Parse(dest)
	if err != nil {
		return "", fmt.Errorf("url parse error: %v", err)
	}
	if u.Scheme != "s3" {
		return "", fmt.Errorf("Scheme is not s3")
	}

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 20 * 1024 * 1024 // 10MB part size, default is 5MB
		u.Concurrency = 200           // The number of goroutines when sending parts, default is 5
		u.LeavePartsOnError = true    // Don't delete the parts if the upload fails.

	})
	// Upload input parameters
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(u.Host),
		Key:    aws.String(strings.Trim(u.Path, "/")),
		Body:   f,
	}

	// Perform upload with options different than the those in the Uploader.
	result, err := uploader.Upload(upParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return aws.StringValue(&result.Location), nil
}
