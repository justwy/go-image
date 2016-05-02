package draw_test

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/justwy/treqme/cognitiveservice"
	"github.com/justwy/treqme/draw"
)

func TestDrawRectangleFromInternet(t *testing.T) {

}

func TestDrawRectangleFromReader(t *testing.T) {

}

func ExampleDrawRectangleFromReader() {
	imgFile, err := os.Open("./testdata/white.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, _, _ := image.Decode(imgFile)

	// Draw a rectangle of width 20 and height 10 with the top left point at (10, 20)
	detectInfos := []cognitiveservice.DetectInfo{
		cognitiveservice.DetectInfo{
			FaceRectangle: cognitiveservice.FaceRectangle{
				Width:  10,
				Height: 10,
				Left:   5,
				Top:    5,
			},
			FaceAttributes: cognitiveservice.FaceAttributes{
				Age: 25,
			},
		},
		cognitiveservice.DetectInfo{
			FaceRectangle: cognitiveservice.FaceRectangle{
				Width:  20,
				Height: 20,
				Left:   10,
				Top:    10,
			},
			FaceAttributes: cognitiveservice.FaceAttributes{
				Age: 35,
			},
		},
	}
	processed, _ := draw.DetectInfo(img, detectInfos)
	fmt.Println("debug: ", err)

	f, _ := os.Create("./testdata/rect.jpg")

	buf := bufio.NewWriter(f)
	png.Encode(buf, processed)

	buf.Flush()
	// Output: debug:  <nil>
}
