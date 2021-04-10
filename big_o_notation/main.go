package main

import "fmt"

func main() {
	bigOOne(1)
	bigON([]int{1, 2, 3})
	bigOlogN(10) // 10 5 2
	bigONlogN(10)
	bigON2([]int{1, 2, 3, 4, 5})
	sum(4)
	pairSumSequence(4) // 1 4 9 16
}

// O(1)
func bigOOne(n int) {
	fmt.Println(n) //
}

// bigON O(n)
func bigON(n []int) {
	for i := 0; i < len(n); i++ {
		fmt.Println(i)
	}
}

// O(log(n))
func bigOlogN(n int) {
	if n <= 1 {
		return
	}
	fmt.Println(n)
	bigOlogN(n / 2)
}

// O(n + log(n))
func bigONlogN(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	if n <= 1 {
		return
	}
	bigONlogN(n / 2)
}

// O(n**2)
func bigON2(n []int) {
	for _, i := range n {
		for _, j := range n {
			fmt.Println(i, j)
		}
		fmt.Println()
	}
}

// O(a + b)
func bigOAPlusB(a []int, b []int) {
	for _, i := range a {
		fmt.Println(i)
	}
	for _, i := range b {
		fmt.Println(i)
	}
}

// O(a * b)
func bigOAMultipliedB(a []int, b []int) {
	for _, i := range a {
		for _, n := range b {
			fmt.Println(i)
			fmt.Println(n)
		}
	}
}

func sum(n int) int {
	fmt.Println(n)
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}

func pairSumSequence(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += pairSum(i, i+1)
		fmt.Println(sum)
	}
	return sum
}

func pairSum(a int, b int) int {
	return a + b
}
