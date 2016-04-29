package s3uploader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
	"io"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

type Uploader interface {
	Upload(f io.ReadSeeker, hours int) (url string, err error)
}

type S3Uploader struct {
	//Uploader *s3manager.Uploader
	S3SVC *s3.S3
	S3Bucket string
	S3Key string
}

/*
func (s3Uploader S3Uploader) Upload(f io.Reader) (url string, err error) {
	result, err := s3Uploader.Uploader.Upload(&s3manager.UploadInput{
		Body:   f,
		Key:    aws.String(s3Uploader.S3Key),
		Bucket: &s3Uploader.S3Bucket,
	})

	if err != nil {
		return "", errors.Wrap(err, "Failed to upload image to S3 " + s3Uploader.S3Key)
	}

	return result.Location, nil
}
*/

func (s3Uploader S3Uploader) Upload(f io.ReadSeeker, hours int) (url string, err error) {
	_, err = s3Uploader.S3SVC.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Uploader.S3Bucket),
		Key: aws.String(s3Uploader.S3Key),
		Body: f,
	})

	if err != nil {
		return url, errors.Wrap(err, "Failed to upload object to S3")
	}

	req, _ := s3Uploader.S3SVC.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3Uploader.S3Bucket),
		Key: aws.String(s3Uploader.S3Key),
	})

	urlStr, err := req.Presign(1 * time.Hour)
	if err != nil {
		return urlStr, errors.Wrap(err, "Failed to sign request")
	}

	return urlStr, nil
}

