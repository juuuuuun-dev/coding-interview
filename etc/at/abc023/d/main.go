package main

import (
	"fmt"
	"sort"
)

func calc(N int, h, s []int) {
	left := 0
	right := 100000
	t := make([]int, N)
	for right-left > 1 {
		mid := (left + right) / 2
		fmt.Println("mid", mid)
		ok := true
		for i := 0; i < N; i++ {
			if mid < h[i] {
				ok = false
			} else {
				t[i] = (mid - h[i]) / s[i]
			}
		}
		sort.Ints(t)
		fmt.Println(t)
		for i := 0; i < N; i++ {
			if t[i] < i {
				ok = false
			}
		}
		fmt.Println("ok", ok)
		if ok {
			right = mid
		} else {
			left = mid
		}
	}
	fmt.Println(right)
}

func main() {
	var N int
	fmt.Scan(&N)
	h := make([]int, N)
	s := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &h[i], &s[i])
	}
	calc(N, h, s)
}
