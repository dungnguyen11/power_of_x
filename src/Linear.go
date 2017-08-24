package main

import (
	"fmt"
	"TimeTrack"
	"time"
)

func main() {
	defer TimeTrack.TimeTrack(time.Now(), "Iterative")

	var x float64 = 2
	var n int = 15

	var result = recursiveLinear(x, n)
	var result2 = iterativeLinear(x, n)
	fmt.Println("Recursive result", x, "power", n, "=", result)
	fmt.Println("Iterative result", x, "power", n, "=", result2)
}


func recursiveLinear(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	return x * recursiveLinear(x, n - 1)
}


func iterativeLinear(x float64, n int) float64 {
	var result float64 = 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}


