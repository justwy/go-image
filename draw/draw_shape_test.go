package draw_test

import (
	"testing"
	"fmt"
	"github.com/justwy/treqme/draw"
	"os"
	"log"
	"image"
	"bufio"
	"image/png"
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
	processed, _ := draw.DrawRectangle(img, []image.Rectangle{image.Rect(5, 5, 10, 10), image.Rect(10, 10, 20, 20)})
	fmt.Println("debug: ", err)

	f, _ := os.Create("./testdata/rect.jpg")

	buf := bufio.NewWriter(f)
	png.Encode(buf, processed)

	buf.Flush()
	// Output: debug:  <nil>
}