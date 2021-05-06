package main

import (
	"fmt"
)

func calc(n int, nums []int, count int) int {
	flg := true
	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			nums[i] = nums[i] / 2
		} else {
			flg = false
			break
		}
	}
	if flg {
		return calc(n, nums, count+1)
	}
	return count
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	nums := make([]int, N)
	fmt.Scanf("%d %d %d", &nums[0], &nums[1], &nums[2])
	c := calc(N, nums, 0)
	fmt.Println(c)
}
