// Package draw provides functions to draw basic shapes and text on image.
package draw

import (
	"image"
	"image/color"
	"image/draw"
	"strconv"

	"github.com/justwy/treqme/cognitiveservice"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/pkg/errors"
)

func init() {
	draw2d.SetFontFolder("./resource/font/")
}

// DetectInfo draws rectangles with given sizes.
func DetectInfo(srcImg image.Image, detectInfos []cognitiveservice.DetectInfo) (image.Image, error) {
	// create a copy of srcImg to dstImg
	r := srcImg.Bounds()
	dstImg := image.NewRGBA(srcImg.Bounds())
	draw.Draw(dstImg, r, srcImg, r.Min, draw.Src)

	// convert detectInfos to rect
	var rects []image.Rectangle
	var ages []string
	for _, face := range detectInfos {
		rects = append(rects, image.Rect(
			int(face.FaceRectangle.Left),
			int(face.FaceRectangle.Top),
			int(face.FaceRectangle.Left+face.FaceRectangle.Width),
			int(face.FaceRectangle.Top+face.FaceRectangle.Height)))
		ages = append(ages, strconv.Itoa(int(face.FaceAttributes.Age)))
	}

	for i, rect := range rects {
		err := drawFace(dstImg, ages[i], float64(rect.Min.X), float64(rect.Dx()), float64(rect.Min.Y), float64(rect.Dy()))
		if err != nil {
			return dstImg, errors.Wrap(err, "Got error while drawing rectangle "+rect.String())
		}
	}

	return dstImg, nil
}

func drawFace(img draw.Image, age string, x, xLen, y, yLen float64) error {

	// Initialize the graphic context on an RGBA image
	gc := draw2dimg.NewGraphicContext(img)

	// Set some properties
	gc.SetFillColor(color.Transparent)
	gc.SetStrokeColor(color.RGBA{0xff, 0x75, 0x1a, 0xff})
	gc.SetLineWidth(2)

	// Draw a closed shape
	gc.MoveTo(x, y) // should always be called first for a new path
	gc.LineTo(x, y+yLen)
	gc.LineTo(x+xLen, y+yLen)
	gc.LineTo(x+xLen, y)
	gc.LineTo(x, y)

	// Draw text
	gc.SetFontData(draw2d.FontData{Name: "luxi", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontSize(30)
	gc.StrokeStringAt("Age: "+age, x, y-5)

	gc.Close()
	gc.FillStroke()

	gc.Restore()

	return nil
}
