package main

import "fmt"

func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func calc(N, A, B int) {
	total := 0
	for i := 0; i <= N; i++ {
		sum := findSumOfDigits(i)
		if sum >= A && sum <= B {
			total += i
		}
	}
	fmt.Println(total)
}

func main() {
	var N, A, B int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	calc(N, A, B)
	fmt.Println(findSumOfDigits(823))
}
