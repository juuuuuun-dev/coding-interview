package main

import "fmt"

func GCD(m, n int) int {
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}
func main() {
	a := GCD(51, 15)
	fmt.Println(a)
	fmt.Println(GCD(15, 51))
}
