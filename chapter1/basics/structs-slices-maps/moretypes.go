package main

import "fmt"

// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

// struct literal denotes a newly allocated struct value by listing the values of the fields
// can list just a subset by using Name: syntax, order is irrelevant
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y: 0 is implicit
	v3 = Vertex{}      // X: 0 and Y: 0 is implicit
	r  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// structs
	fmt.Println(Vertex{1, 2})

	// struct fields accessed using a dot
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// pointers to structs
	q := &v
	q.X = 1e9 // dereference is implicit for struct pointers, no need for (*q).X
	fmt.Println(v)

	// struct literals
	fmt.Println(v1, r, v2, v3)

	// Arrays
	// the type [n]T is an array of n values of type T
	// var a [10]int is an array of ten integers
	// arrays cannot be resized as length is part of it's type
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	// dynamically sized, []T is a slice with elements of type T
	// specify two bounds a[low: high]
	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are like references to arrays
	// changing the elements of a slice modifies the corresponding array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	c[0] = "XXX"
	fmt.Println(b, c)
	fmt.Println(names)

	// slice literals
	d := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(d)

	e := []bool{true, false, true, true, false, true}
	fmt.Println(e)

	f := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(f)
}
