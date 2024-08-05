package canvas

// PixelsToPoints converts pixels to points
// Why Do I need to do this?
// Why does the library multiply the points by 0.352778?
func PixelsToPoints(point float64) float64 {
	return point / 0.352778
}
