package main

import "fmt"

func fibo(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	return fibo(N-1) + fibo(N-2)
}

func fiboDetail(n int) int {
	fmt.Printf("fibo(%d)を呼び出しました\n", n)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	res := fiboDetail(n-1) + fiboDetail(n-2)
	fmt.Printf("ここまでの合計 %d\n", res)
	return res
}

func fiboDP(n int, dp []int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	if dp[n-1] != 0 {
		fmt.Println("cache", dp[n-1])
		return dp[n-1]
	}
	dp[n-1] = fiboDP(n-1, dp) + fiboDP(n-2, dp)
	fmt.Println("ここまでの合計", dp)
	return dp[n-1]
}

func main() {
	// fmt.Println(fibo(4))
	// fmt.Println(fiboDetail(6))
	t := 7
	dp := make([]int, t)
	fmt.Println(fiboDP(t, dp))
}
