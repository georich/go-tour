package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, only in loop scope
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post statements are optional
	sumTwo := 1
	for sumTwo < 1000 {
		sumTwo += sumTwo
	}
	fmt.Println(sumTwo)

	// for is Go's while
	sumThree := 1
	for sumThree < 1000 {
		sumThree += sumThree
	}
	fmt.Println(sumThree)

	// infinite loop
	// for {
	// }

	// if loop
	fmt.Println(sqrt(2), sqrt(-4))

	// if with short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
