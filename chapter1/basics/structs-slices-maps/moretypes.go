package main

import (
	"fmt"
	"strings"
)

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

	// slice defaults
	// default lower bound is 0 and default higher bound is length of slice
	z := []int{2, 3, 5, 7, 11, 13}

	z = z[1:4]
	fmt.Println(z)

	z = z[:2]
	fmt.Println(z)

	z = z[1:]
	fmt.Println(z)

	// slice length and capacity
	// length is number of elements
	// capacity is number of elements in underlying array, counting from first element of slice
	n := []int{2, 3, 5, 7, 11, 13}
	printSlice(n)

	// slice the slice to give it zero length
	n = n[:0]
	printSlice(n)

	// extend its length
	n = n[:4]
	printSlice(n)

	// drop its first two values
	n = n[2:]
	printSlice(n)

	// nil slices
	// zero value of a slice is nil
	var o []int
	fmt.Println(o, len(o), cap(o))
	if o == nil {
		fmt.Println("nil!")
	}

	// creating a slice with make
	t := make([]int, 5)
	printSlice(t)

	t2 := make([]int, 0, 5)
	printSlice(t2)

	t3 := t2[:2]
	printSlice(t3)

	t4 := t3[2:5]
	printSlice(t4)

	// slices of slices
	// create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// appending to a slice
	var h []int
	printSlice(h)

	h = append(h, 0)
	printSlice(h)

	h = append(h, 1)
	printSlice(h)

	h = append(h, 2, 3, 4)
	printSlice(h)

	// Range
	// range form of the for loop iterates over a slice or map
	// ranging over a slice returns two values, first is index, second is copy of element value
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// can skip index or value by assigning to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// Maps
	// a map maps keys to vales
	var m map[string]VertexTwo

	m = make(map[string]VertexTwo)
	m["Bell Labs"] = VertexTwo{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	// map literals
	var m2 = map[string]VertexTwo{
		"Bell Labs": VertexTwo{
			40.68433, -74.39967,
		},
		"Google": VertexTwo{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// if the top-level type is just a name it can be omitted from the elements of the literal
	var m3 = map[string]VertexTwo{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	// mutating Maps
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	value, ok := m["Answer"]
	fmt.Println("The value:", value, "Present?", ok)
}

type VertexTwo struct {
	Lat, Long float64
}

func printSlice(n []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(n), cap(n), n)
}
