package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &A[i], &B[i])
	}
	sum := 0
	for i := 0; i < N; i++ {
		A[i] += sum
		remainder := A[i] % B[i]
		if remainder != 0 {
			sum += B[i] - remainder
		}
	}
	fmt.Println(sum)
}
