package s3uploader_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/justwy/treqme/face/s3uploader"
)

func ExampleUploadImage() {
	imgUrl := "https://golang.org/doc/gopher/project.png"

	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))

	s3Uploader := s3uploader.S3Uploader{
		S3SVC:    svc,
		S3Bucket: "test-bucket",
		S3Key:    "test/",
	}

	resp, _ := http.Get(imgUrl)
	defer resp.Body.Close()

	bodyInByte, _ := ioutil.ReadAll(resp.Body)

	url, _ := s3Uploader.Upload("id", strings.NewReader(string(bodyInByte)), 2)
	fmt.Println(url)
}
