package handler_test

import (
	"github.com/justwy/treqme/handler"
	"github.com/justwy/treqme/cognitiveservice"
	"github.com/justwy/treqme/s3uploader"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"fmt"
)

func ExampleDetectFace() {
	faceAPI := cognitiveservice.NewMicrosoftFaceAPI("api-key")

	config := aws.NewConfig().WithRegion("us-east-1").WithCredentialsChainVerboseErrors(true)
	sess := session.New(config)
	uploader := s3manager.NewUploader(sess)

	s3Uploader := s3uploader.S3Uploader{
		uploader,
		"treqme",
		"testabc",
	}

	inputUrl := "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Gilbert_Stuart_Williamstown_Portrait_"
	+ "of_George_Washington.jpg/197px-Gilbert_Stuart_Williamstown_Portrait_of_George_Washington.jpg"

	url, err := handler.DetectFace(inputUrl, faceAPI, s3Uploader)
	fmt.Println(url, err)
}

