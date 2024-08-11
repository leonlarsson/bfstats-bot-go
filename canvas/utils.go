package canvas

import (
	"bytes"
	"image"
	"image/png"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

// PixelsToPoints converts pixels to points
// Why Do I need to do this?
// Why does the library multiply the points by 0.352778?
func PixelsToPoints(point float64) float64 {
	return point / 0.352778
}

// CanvasToImage converts a canvas to an image.Image
func CanvasToImage(c *canvas.Canvas) image.Image {
	return rasterizer.Draw(c, canvas.Resolution(1), nil)
}

// CanvasToBuffer converts a canvas to a bytes.Buffer. Used in when sending a File to Discord
func CanvasToBuffer(c *canvas.Canvas) *bytes.Buffer {
	buf := new(bytes.Buffer)
	png.Encode(buf, CanvasToImage(c))
	return buf
}
