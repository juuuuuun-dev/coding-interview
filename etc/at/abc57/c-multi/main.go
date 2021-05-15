package main

import (
	"fmt"
	"math"
	"strconv"
)

func f(n, i int) int {
	return max(len(intToStr(n)), len(intToStr(i)))
}
func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func min(n, i int) int {
	return int(math.Min(float64(n), float64(i)))
}

func intToStr(i int) string {
	var strVal = strconv.Itoa(i)
	return strVal
}

func main() {
	var N, a int
	a = 100
	fmt.Scanf("%d", &N)
	for i := 1; i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		a = min(a, f(i, N/i))
	}
	fmt.Println(a)
}
