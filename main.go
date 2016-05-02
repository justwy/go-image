package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/justwy/treqme/cognitiveservice"
	"github.com/justwy/treqme/handler"
	"github.com/justwy/treqme/s3uploader"
)

var (
	faceAPI = cognitiveservice.NewMicrosoftFaceAPI("131b5264f0954b608d41daac603276cd")

	svc = s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))

	s3Uploader = s3uploader.S3Uploader{
		S3SVC:    svc,
		S3Bucket: "treqme",
		S3Key:    "img/",
	}
)

func processImage(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	imgURL := queryMap["url"][0]
	url, err := handler.DetectFace(imgURL, faceAPI, s3Uploader)
	if err != nil {
		http.Error(w, "Failed to process image with error", http.StatusInternalServerError)
		log.Print(err)
	} else {
		fmt.Fprint(w, url)
	}
}

func main() {
	http.HandleFunc("/detect/", processImage)
	err := http.ListenAndServe(":5100", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
