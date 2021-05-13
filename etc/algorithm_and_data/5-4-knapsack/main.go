package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func split(s string) []string {
	return strings.Fields(s)
}

func strAryToIntAry(strs []string) []int {
	var ret = make([]int, 0, len(strs))
	for _, str := range strs {
		var intVal, e = strconv.Atoi(string(str))
		if e != nil {
			panic(e)
		}
		ret = append(ret, intVal)
	}
	return ret
}

func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func calc(N int, W int, weight []int, value []int) {
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		fmt.Println("i", i)
		for w := 1; w <= W; w++ {
			// 仮で今までの計算を代入
			dp[i+1][w] = dp[i][w]
			// 重みがmaxより低ければ前回までのmax valueと比較してmaxを代入
			// つまり値を更新しながら進める
			if w-weight[i] >= 0 { // w >= weight[i]
				dp[i+1][w] = max(dp[i+1][w], dp[i][w-weight[i]]+value[i])
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(dp[N][W])
}

func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)
	weight := make([]int, N)
	value := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &weight[i], &value[i])
	}
	calc(N, W, weight, value)
}
