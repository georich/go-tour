package main

import "fmt"

// func add(x int, y int) int {
// consecutive type params can be shortened by inserting after the last one
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// return values may be named, will be treated as being defined at top of function
// should document meaning, naked return w/o args will return named values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
