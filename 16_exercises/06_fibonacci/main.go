package main

import "fmt"

func main() {
	fib := fibonacci([]int{1, 2}, 4000000)
	fmt.Println(fib)
	var sumFibEvens int
	for _, v := range fib {
		if v%2 == 0 {
			sumFibEvens += v
		}
	}
	fmt.Println(sumFibEvens)
}

func fibonacci(series []int, limit int) []int {
	sum := series[len(series)-1] + series[len(series)-2]
	if sum < limit {
		return fibonacci(append(series, sum), limit)
	}
	return series
}
