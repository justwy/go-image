package handler

import (
	"net/http"
	"github.com/pkg/errors"
	"github.com/justwy/treqme/cognitiveservice"
	"github.com/justwy/treqme/draw"
	"image"
	"github.com/justwy/treqme/s3uploader"
	"image/png"
	"bytes"
	"strings"
)

func DetectFace(url string, faceAPI cognitiveservice.FaceAPI, uploader s3uploader.Uploader) (returnedURL string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to download image from " + url)
	}

	defer resp.Body.Close()

	faceResult, err := faceAPI.Detect(url)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to detect face " + url)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to decode image at " + url)
	}

	var rects []image.Rectangle
	for _, face := range faceResult {
		rects = append(rects, image.Rect(
			int(face.FaceRectangle.Left),
			int(face.FaceRectangle.Top),
			int(face.FaceRectangle.Left + face.FaceRectangle.Width),
			int(face.FaceRectangle.Top + face.FaceRectangle.Height)))
	}

	//draw.DrawRectangle(img, faceResult)
	processed, err := draw.DrawRectangle(img, rects)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to decode image at " + url)
	}

	buf := new(bytes.Buffer)
	png.Encode(buf, processed)

	returnedURL, err = uploader.Upload(strings.NewReader(string(buf.Bytes())), 1)
	if err != nil {
		return returnedURL, errors.Wrap(err, "Failed to upload processed image to S3")
	}
	return returnedURL, nil
}
