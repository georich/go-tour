package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	// Pi is an exported name and must start with a capital letter
	fmt.Println(math.Pi)
}
