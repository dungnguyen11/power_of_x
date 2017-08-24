package main

import (
	"TimeTrack"
	"Exe1.1"
	"time"
	"fmt"
)

func main() {
	defer TimeTrack.TimeTrack(time.Now(), "Linear")


	var x float64 = 2
	var n float64 = 10000

	//Using recursive
	//var result1 float64 = recursiveLinear(x, n)
	//fmt.Println("Recursive result = ", result1)

	//Using iterative
	var result2 float64 = iterativeLinear(x, n)
	fmt.Println("Iteration result = ", result2)

		//defer timeTrack(time.Now(), "Logarithmic")
		//var x float64 = 2
		//var n int = 4

		//Using recursive
		//var result float64 = recursiveLogarithmic(x, n)
		//fmt.Println("Recursive result: ", result)

		//Using iterative
		var result float64 = iterativeLogarithmic(x, n)
		fmt.Println("Iterative result: ", result)

	}