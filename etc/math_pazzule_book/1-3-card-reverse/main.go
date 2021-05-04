package main

import "fmt"

func calc(n int) {
	cards := make([]bool, n)
	for i := 2; i <= n; i++ {
		j := i - 1
		for j < n {
			cards[j] = !cards[j]
			j += i
		}
	}
	fmt.Println(cards)
	for i := 0; i < n; i++ {
		if !cards[i] {
			fmt.Println(i + 1)
		}
	}
}

func calc2(n int) {
	fmt.Println(n)
	for i := 1; i <= n; i++ {
		// fmt.Println(i)
		flag := false
		for j := 1; j <= n; j++ {
			if i%j == 0 {
				flag = !flag
			}
		}
		if flag {
			fmt.Println(i)
		}

	}
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	calc(n)
	// calc2(n)
}
