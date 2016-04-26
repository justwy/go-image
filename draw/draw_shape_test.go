package draw_test

import (
	"testing"
	"fmt"
	"github.com/justwy/treqme/draw"
	"os"
	"log"
)

func TestDrawRectangleFromInternet(t *testing.T) {

}

func TestDrawRectangleFromReader(t *testing.T) {

}

func ExampleDrawRectangleFromReader() {
	imgFile, err := os.Open("../testdata/white.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Draw a rectangle of width 20 and height 10 with the top left point at (10, 20)
	err = draw.DrawRectangleFromReader(imgFile, "../testdata/white-modified.png", 10, 20, 20, 10)
	fmt.Println("debug: ", err)
	// Output: debug:  <nil>
}