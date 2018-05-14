package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		change := (z*z - x) / (2 * z)
		if math.Abs(change) < 0.001 {
			return z
		}
		z -= change
		fmt.Println(z)
	}
	return z
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

	// sqrt exercise
	fmt.Println(Sqrt(49))

	// switch statement
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch with no condition is equivalent to switch true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
