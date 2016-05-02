package handler_test

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/justwy/treqme/cognitiveservice"
	"github.com/justwy/treqme/handler"
	"github.com/justwy/treqme/s3uploader"
)

func ExampleDetectFace() {
	faceAPI := cognitiveservice.NewMicrosoftFaceAPI("api-key")

	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))

	s3Uploader := s3uploader.S3Uploader{
		S3SVC:    svc,
		S3Bucket: "treqme",
		S3Key:    "testabc",
	}

	inputUrl := "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Gilbert_Stuart_Williamstown_Portrait_of_George_Washington.jpg/197px-Gilbert_Stuart_Williamstown_Portrait_of_George_Washington.jpg"

	url, err := handler.DetectFace(inputUrl, faceAPI, s3Uploader)
	fmt.Println(url, err)
}
