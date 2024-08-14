package canvas

import (
	"github.com/tdewolff/canvas"
)

// CreateCanvasAndContext creates a canvas and a context
func CreateStatsCanvasAndContext() (*canvas.Canvas, *canvas.Context) {
	c := canvas.New(1200, 750)
	ctx := canvas.NewContext(c)

	// Set the coordinate system to match what I am used to: (0, 0) in the top left corner
	ctx.SetCoordSystem(canvas.CartesianIV)
	return c, ctx
}

// CreateCanvas creates a canvas
func CreateCanvas(width float64, height float64) *canvas.Canvas {
	return canvas.New(width, height)
}

// CreateContext creates a context
func CreateContext(c *canvas.Canvas) *canvas.Context {
	return canvas.NewContext(c)
}
