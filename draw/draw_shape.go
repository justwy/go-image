// Package draw provides functions to draw basic shapes and text on image.
package draw

import (
	"image"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"image/draw"
	"github.com/pkg/errors"
	_ "image/jpeg"
	_ "image/png"
)

// DrawRectangle draws rectangles with given sizes.
func DrawRectangle(srcImg image.Image, rectangles []image.Rectangle) (image.Image, error) {
	// create a copy of srcImg to dstImg
	r := srcImg.Bounds()
	dstImg := image.NewRGBA(srcImg.Bounds())
	draw.Draw(dstImg, r, srcImg, r.Min, draw.Src)

	for _, rect := range rectangles {
		err := drawRectangle(dstImg, float64(rect.Min.X), float64(rect.Dx()), float64(rect.Min.Y), float64(rect.Dy()))
		if err != nil {
			return dstImg, errors.Wrap(err, "Got error while drawing rectangle " + rect.String())
		}
	}

	return dstImg, nil
}

func drawRectangle(img draw.Image, x, xLen, y, yLen float64) error {

	// Initialize the graphic context on an RGBA image
	gc := draw2dimg.NewGraphicContext(img)

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

	return nil
}