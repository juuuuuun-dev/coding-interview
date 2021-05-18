package main

import (
	"fmt"
	"math"
)

func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)
	w := make([]int, N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &w[i], &v[i])
	}
	fmt.Println(w, v)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i < N; i++ {
		for j := 1; j <= W; j++ {
			dp[i+1][j] = dp[i][j]
			if W-w[i] >= 0 {
				fmt.Println("dp[i+1][W-w[i]]", dp[i+1][W-w[i]])
				dp[i+1][j] = max(dp[i+1][j], dp[i+1][W-w[i]]+w[i])
			}
		}
	}
	fmt.Println(dp)

}
