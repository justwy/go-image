package handler

import (
	"bytes"
	"image"
	"image/png"
	"net/http"
	"strings"

	"github.com/justwy/treqme/face/cognitiveservice"
	"github.com/justwy/treqme/face/draw"
	"github.com/justwy/treqme/face/s3uploader"
	"github.com/pkg/errors"
)

// DetectFace detects faces from url using microsoft coginitive service API.
func DetectFace(url string, faceAPI cognitiveservice.FaceAPI, uploader s3uploader.Uploader) (returnedURL string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to download image from "+url)
	}

	defer resp.Body.Close()

	detectInfos, err := faceAPI.Detect(url)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to detect face "+url)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to decode image at "+url)
	}

	//draw.DrawRectangle(img, faceResult)
	processed, err := draw.DetectInfo(img, detectInfos)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to decode image at "+url)
	}

	buf := new(bytes.Buffer)
	png.Encode(buf, processed)

	returnedURL, err = uploader.Upload(url, strings.NewReader(string(buf.Bytes())), 1)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to upload processed image to S3")
	}
	return returnedURL, nil
}
