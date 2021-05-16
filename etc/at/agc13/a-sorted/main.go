package main

import (
	"bufio"
	"fmt"
	"os"
)

func binarySearch(x int, nums []int) int {
	l := 0
	r := len(nums)
	for r-1 >= l {
		mid := (l + r) / 2
		if nums[mid] == x {
			return mid
		} else if nums[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}

func rec(N int, A []int, c [][]int) [][]int {
	l := len(A)
	if l == 0 {
		return c
	}
	fmt.Println("A", A)
	temp := []int{}
	v := A[0]
	temp = append(temp, v)
	if l > 1 {
		A = A[1:]
		A, temp = rec2(v, A, temp)
		c = append(c, temp)
		return rec(N, A, c)
	}
	c = append(c, temp)
	return c
}

func rec2(x int, A []int, nums []int) ([]int, []int) {
	index := binarySearch(x+1, A)
	if index == -1 {
		return A, nums
	}
	nums = append(nums, A[index])
	A = append(A[:index], A[index+1:]...)
	return rec2(x+1, A, nums)
}

func calc(N int, A []int) {
	var op int
	ans := 1
	start := A[0]
	for _, v := range A {
		switch op {
		case 0:
			if v > start {
				op = 1
			} else {
				op = -1
			}
		case 1:
			if v < start {
				op = 0
				ans++
			}
		case -1:
			if v > start {
				op = 0
				ans++
			}
		}
		start = v
	}
	fmt.Println(ans)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &A[i])
	}
	calc(N, A)
	// sort.Ints(A)
	// c := [][]int{}
	// fmt.Println(A)
	// res := rec(N, A, c)
	// fmt.Println(res)

}
