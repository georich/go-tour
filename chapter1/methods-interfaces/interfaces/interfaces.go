package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

// Interfaces are implemented implicitly
type I interface {
	M()
}

type T struct {
	S string
}

// method means type T implements interface I, don't have to explicitly declare this
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	// interfaces implicitly
	// var i I = T{"hello"}
	// i.M()

	// interface values
	var i2 I

	i2 = &T{"Hello"}
	describe(i2)
	i2.M()

	i2 = F(math.Pi)
	describe(i2)
	i2.M()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// interface values
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
