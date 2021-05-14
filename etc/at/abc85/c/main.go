package main

import "fmt"

func calc(N, Y int) (int, int, int) {
	res10000, res5000, res1000 := -1, -1, -1
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			k := N - i - j
			if k < 0 {
				break
			}
			if 10000*i+5000*j+1000*k == Y {
				res10000 = i
				res5000 = j
				res1000 = k
				break
			}
		}
	}
	return res10000, res5000, res1000
}

func main() {
	var N, Y int
	fmt.Scanf("%d %d", &N, &Y)
	a, b, c := calc(N, Y)
	fmt.Println(a, b, c)
}
