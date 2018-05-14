package main

import "fmt"

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
}
