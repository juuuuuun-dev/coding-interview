package main

import (
	"fmt"
	"math"
)

func min(n, i int) int {
	return int(math.Min(float64(n), float64(i)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func main() {
	var A, B, C, X, Y int
	fmt.Scanf("%d %d %d %d %d", &A, &B, &C, &X, &Y)
	ab := C * 2
	total := 0
	if A+B > ab {
		minN := min(X, Y)
		maxN := max(X, Y)
		sub := abs(X - Y)
		total += ab * minN
		if X > Y {
			total += A * sub
		} else {
			total += B * sub
		}
		total = min(total, maxN*ab)
	} else {
		total += A*X + B*Y
	}
	fmt.Println(total)
}
