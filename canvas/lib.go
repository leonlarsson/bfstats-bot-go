package canvas

import (
	"github.com/tdewolff/canvas"
)

// CreateCanvasAndContext creates a canvas and a context
func CreateCanvasAndContext(width float64, height float64) (*canvas.Canvas, *canvas.Context) {
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)
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
