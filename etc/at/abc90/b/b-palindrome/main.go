package main

import "fmt"

func reversInt(n int) int {
	res := 0
	for n > 0 {
		remainder := n % 10
		res *= 10
		res += remainder
		n /= 10
	}
	return res
}

func calc(A, B int) {
	c := 0
	for i := A; i <= B; i++ {
		num := reversInt(i)
		if i == num {
			c += 1
			// fmt.Println(i)
		}
	}
	fmt.Println(c)
}

func main() {
	var A, B int
	fmt.Scanf("%d %d", &A, &B)
	calc(A, B)
}
