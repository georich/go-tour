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
}
