// Package draw provides functions to draw basic shapes and text on image.
package draw

import (
	"io"
	"image"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"net/http"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
)

func DrawRectangleFromInternet(url string, dstPath string, x, xLen, y, yLen float64) error {
	resp, err := http.Get(url)

	if err != nil {
		return err;
	}

	defer resp.Body.Close()

	return DrawRectangleFromReader(resp.Body, dstPath, x, xLen, y, yLen)
}

func DrawRectangleFromReader(sourceImg io.Reader, dstPath string, x, xLen, y, yLen float64) error {
	img, _, err := image.Decode(sourceImg);
	if err != nil {
		return err
	}

	processedImg, err := drawRectangle(img, x, xLen, y, yLen)

	return draw2dimg.SaveToPngFile(dstPath, processedImg)

	return err
}

func drawRectangle(srcImg image.Image, x, xLen, y, yLen float64) (image.Image, error) {
	// create a copy of srcImg to dstImg
	r := srcImg.Bounds()
	dstImg := image.NewRGBA(srcImg.Bounds())
	draw.Draw(dstImg, r, srcImg, r.Min, draw.Src)

	// Initialize the graphic context on an RGBA image
	gc := draw2dimg.NewGraphicContext(dstImg)

	// Set some properties
	gc.SetFillColor(color.Transparent)
	gc.SetStrokeColor(color.RGBA{0xff, 0x75, 0x1a, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(x, y) // should always be called first for a new path
	gc.LineTo(x, y + yLen)
	gc.LineTo(x + xLen, y + yLen)
	gc.LineTo(x + xLen, y)
	gc.LineTo(x, y)
	gc.Close()
	gc.FillStroke()

	gc.Restore()

	return dstImg, nil
}