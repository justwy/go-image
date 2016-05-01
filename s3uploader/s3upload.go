package s3uploader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
	"io"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
	"encoding/base64"
	"math/rand"
	"fmt"
	"strconv"
)

type Uploader interface {
	Upload(id string, f io.ReadSeeker, hours int) (url string, err error)
}

type S3Uploader struct {
	S3SVC *s3.S3
	S3Bucket string
	S3Key string
}

func init() {
	seed :=	time.Now().UTC().UnixNano()
	rand.Seed(seed)
	fmt.Println("seed: ", seed)
}

func (s3Uploader S3Uploader) Upload(id string, f io.ReadSeeker, hours int) (url string, err error) {
	fullKey := aws.String(s3Uploader.S3Key + strconv.Itoa(rand.Intn(999999)) + "_" +  base64.StdEncoding.EncodeToString([]byte(id)))
	_, err = s3Uploader.S3SVC.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Uploader.S3Bucket),
		Key: fullKey,
		Body: f,
	})

	if err != nil {
		return url, errors.Wrap(err, "Failed to upload object to S3")
	}

	req, _ := s3Uploader.S3SVC.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Uploader.S3Bucket),
		Key: fullKey,
	})

	urlStr, err := req.Presign(24 * time.Hour)
	if err != nil {
		return urlStr, errors.Wrap(err, "Failed to sign request")
	}

	return urlStr, nil
}

