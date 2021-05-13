package main

import "fmt"

var cnt = 0

func change(n int, coins []int, max int, count int) {
	coin := coins[0]
	coins = coins[1:]
	if len(coins) == 0 {
		if n/coin <= max {
			fmt.Println("n", n)
			fmt.Println("coin", coin)
			cnt += 1
		}
	} else {
		for j := 0; j <= n/coin; j++ {
			change(n-coin*j, coins, max-j, count)
		}
	}
}

func roop() {
	c := 0
	for c500 := 0; c500 <= 2; c500++ {
		for c100 := 0; c100 <= 10; c100++ {
			for c50 := 0; c50 <= 15; c50++ {
				for c10 := 0; c10 <= 10; c10++ {
					if c500+c100+c50+c10 <= 15 && c500*500+c100*100+c50*50+c10*10 == 1000 {
						c += 1
					}
				}
			}
		}
	}
	fmt.Println("c", c)
}

func main() {
	n := 1000
	coins := []int{500, 100, 50, 10}
	max := 15
	change(n, coins, max, 0)
	fmt.Println("cnt", cnt)
	roop()
}
