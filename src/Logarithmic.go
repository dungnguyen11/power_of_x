package main

import (
	"fmt"
	"time"
	"TimeTrack"
)

func main() {
	defer TimeTrack.TimeTrack(time.Now(), "Logarithmic")

	var x float64 = 3
	var n int = 7

	//Using recursive
	var result = recursiveLogarithmic(x, n)
	fmt.Println("Recursive result = ", result)

	//Using iterative
	var result1 = iterativeLogarithmic(x, n)
	fmt.Println("Iterative result = ", result1)
}

func recursiveLogarithmic(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n % 2 == 0 {
		return recursiveLogarithmic(x, n/2) * recursiveLogarithmic(x, n/2)
	} else {
		return  x * recursiveLogarithmic(x, n/2) * recursiveLogarithmic(x, n/2)
	}
}

func iterativeLogarithmic(x float64, n int) float64 {
	var result float64 = x
	var extra float64 = 1

	if n == 0 {
		return 1
	}

	for n > 1 {
		if n % 2 != 0 {
			extra *= result
			n--
		}
		result *= result
		n = n / 2
	}
	return result * extra
}

