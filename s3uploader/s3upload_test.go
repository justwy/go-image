package s3uploader_test

import (
	"github.com/justwy/treqme/s3uploader"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"net/http"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func ExampleUploadImage() {
	imgUrl := "https://golang.org/doc/gopher/project.png"

	resp, _ := http.Get(imgUrl)
	defer resp.Body.Close()

	config := aws.NewConfig().WithRegion("us-east-1").WithCredentialsChainVerboseErrors(true)
	sess := session.New(config)
	uploader := s3manager.NewUploader(sess)
	url, _ := s3uploader.Upload(resp.Body, uploader, "treqme", "test/go")
	fmt.Println(url)
}


