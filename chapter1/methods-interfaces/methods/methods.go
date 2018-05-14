package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// has a receiver of type Vertex named v
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// can declare methods on non-struct types too
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers
// methods with pointer receievers can modify the value the receiver points to

// Scale multiplies the struct fields by a factor of f
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// pointer receivers
	ve := Vertex{3, 4}
	ve.Scale(10) // interpreted as (&ve).Scale(10) for niceness
	fmt.Println(ve.Abs())
}
