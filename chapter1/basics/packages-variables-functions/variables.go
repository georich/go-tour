package main

import "fmt"

// var c, python, java bool

// variables can take initializers
var i, j int = 1, 2

func main() {
	// var i int
	// var c, python, java = true, false, "no!"

	// short variable declaration - only in functions
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
