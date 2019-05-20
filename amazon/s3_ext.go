package amazon

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/luopengift/log"
)

type customReader struct {
	fp   *os.File
	size int64
	read int64
}

func (r *customReader) Read(p []byte) (int, error) {
	return r.fp.Read(p)
}

func (r *customReader) ReadAt(p []byte, off int64) (int, error) {
	n, err := r.fp.ReadAt(p, off)
	if err != nil {
		return n, err
	}

	// Got the length have read( or means has uploaded), and you can construct your message
	atomic.AddInt64(&r.read, int64(n))

	// I have no idea why the read length need to be div 2,
	// maybe the request read once when Sign and actually send call ReadAt again
	// It works for me
	log.Infof("total read:%d    progress:%d%%\n", r.read/2, int(float32(r.read*100/2)/float32(r.size)))

	return n, err
}

func (r *customReader) Seek(offset int64, whence int) (int64, error) {
	return r.fp.Seek(offset, whence)
}

// UploadExt upload ext
func UploadExt(region, accessKeyID, secretAccessKey, src, dest string) error {
	// The session the S3 Uploader will use

	sess := session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})

	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", src, err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		return fmt.Errorf("file status error: %v", err)
	}

	reader := &customReader{
		fp:   f,
		size: fileInfo.Size(),
	}

	u, err := url.Parse(dest)
	if err != nil {
		return fmt.Errorf("url parse error: %v", err)
	}
	if u.Scheme != "s3" {
		return fmt.Errorf("Scheme is not s3")
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
		Body:   reader,
	}
	fmt.Println("start upload...")
	// Perform upload with options different than the those in the Uploader.
	result, err := uploader.Upload(upParams)
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %#v\n", aws.StringValue(&result.Location))
	return nil
}
