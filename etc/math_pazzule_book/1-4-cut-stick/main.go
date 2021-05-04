package main

import "fmt"

func calc(n int, m int, c int, count int) int {
	if c >= n {
		return count
	} else if c < m {
		return calc(n, m, c*2, count+1)
	} else {
		return calc(n, m, c+m, count+1)
	}
}

func calc2(n int, m int) int {
	count := 0
	current := 1
	for n > current {
		if current < m {
			current += current
		} else {
			current += m
		}
		count++
	}
	return count
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println("calc", calc(n, m, 1, 0))
	fmt.Println("calc2", calc2(n, m))
}
