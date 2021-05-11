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

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func min(n, i int) int {
	return int(math.Min(float64(n), float64(i)))
}

func calc(n int, h []int, dp []int) {
	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = abs(h[i] - h[i-1])
		} else {
			dp[i] = min(dp[i-1]+abs(h[i]-h[i-1]), dp[i-2]+abs(h[i]-h[i-2]))
		}
	}
	fmt.Println(h)
	fmt.Println(dp[n-1])
}

// func chmin(dp *[]int, i int, b int) {
// 	fmt.Println((*dp)[i], i, b)
// 	if (*dp)[i] > b {
// 		(*dp)[i] = b
// 	}
// }

func main() {
	var N int
	fmt.Scanf("%d", &N)
	h := strAryToIntAry(split(readLine()))
	dp := make([]int, N)
	calc(N, h, dp)
}
