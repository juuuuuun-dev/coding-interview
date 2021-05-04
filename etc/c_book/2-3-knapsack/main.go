package main

import (
	"fmt"
	"math"
)

func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)
	w, v := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &w[i], &v[i])
	}
	fmt.Println(w, v)

	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, M+1)
	}
	for i := 1; i <= N; i++ {
		for j := 0; j <= M; j++ {
			ii := i - 1
			dp[i][j] = dp[ii][j]
			if j >= w[ii] {
				dp[i][j] = max(dp[i][j], dp[ii][j-w[ii]]+v[ii])
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(dp[N][M])
}
