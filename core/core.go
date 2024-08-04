package core

import "math"

// Vector is a 3D vector
type Vector struct {
	X, Y, Z float64
}

// NewVector creates a new Vector and returns it
func (a Vector) Add(b Vector) Vector {
	return Vector{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z,
	}
}

// Sub subtracts two vectors and returns a new Vector
func (a Vector) Sub(b Vector) Vector {
	return Vector{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z,
	}
}

// MultiplyByScalar multiplies a Vector by a scalar and returns a new Vector
func (a Vector) MultiplyByScalar(s float64) Vector {
	return Vector{
		a.X * s,
		a.Y * s,
		a.Z * s,
	}
}

// Div divides a Vector by another Vector and returns a new Vector
func (a Vector) Div(b Vector) Vector {
	return Vector{
		a.X / b.X,
		a.Y / b.Y,
		a.Z / b.Z,
	}
}

// Length returns the length of a Vector
func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Dot returns the dot product of two Vectors
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross returns the cross product of two Vectors
func (a Vector) Cross(b Vector) Vector {
	return Vector{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

// Normalize returns the normalized Vector
func (a Vector) Normalize() Vector {
	return a.MultiplyByScalar(1. / a.Length())
}
