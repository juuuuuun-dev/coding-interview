package main

import "fmt"

func rec(n int) int {
	if n == 0 {
		return 0
	}
	fmt.Println("n", n)
	return n + rec(n-1)
}

func rec2(n int) int {
	fmt.Printf("func(%d)を呼び出した\n", n)
	if n == 0 {
		return 0
	}
	res := n + rec2(n-1)
	fmt.Printf("%d までの合計 %d\n", n, res)
	return res
}

func main() {
	// n := rec(5)
	n := rec2(5)
	fmt.Println(n)
}
