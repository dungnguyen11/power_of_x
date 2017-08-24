package main

import (
	"fmt"
	"TimeTrack"
	"time"
)

func main() {
	defer TimeTrack.TimeTrack(time.Now(), "Fibo")

	var n int64 = 35
	var result = recursiveFibo(n)
	var result2 = int64(iterativeFibo(n))
	fmt.Println("recursive Fibo of",n, "=", result)
	fmt.Println("iterative Fibo of",n, "=", result2)
}

func recursiveFibo(n int64) int64{
	var fibo = make([]int64, n + 1)

	if n <= 1 {
		return n
	}

	if fibo[n] != 0 {
		//fmt.Println("1-",fibo[n])
		return fibo[n]
	} else {
		fibo[n] = recursiveFibo(n - 1) + recursiveFibo(n - 2)
		//fmt.Println("n", n)
		//fmt.Println("2-",fibo[n])
		return fibo[n]
	}
}

func iterativeFibo(n int64) int64{
	if n <= 1 {
		return n
	}

	first := int64(0)
	second := int64(1)
	result := int64(0)

	for i := int64(1); i < n;i++ {
		result = first + second
		first = second
		second = result
	}
	return result
}
