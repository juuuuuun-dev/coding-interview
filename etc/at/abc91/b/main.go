package main

import (
	"fmt"
	"math"
)

type Count map[string]int

func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func calc(N, M int, red, blue []string) {
	count_red := count(red, N)
	count_blue := count(blue, M)
	ans := 0
	for k, v := range count_red {
		sub := 0
		if _, ok := count_blue[k]; ok {
			sub = count_red[k] - count_blue[k]
			if sub > 0 {
				ans = max(ans, sub)
			}
		} else {
			ans = max(ans, v)
		}
	}
	fmt.Println(ans)
}

func count(c []string, n int) Count {
	res := make(map[string]int, n)
	for _, v := range c {
		res[v] += 1
	}
	return res
}

func main() {
	var N, M int
	fmt.Scanf("%d", &N)
	red := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &red[i])
	}
	fmt.Scanf("%d", &M)
	blue := make([]string, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%s", &blue[i])
	}
	calc(N, M, red, blue)
}
