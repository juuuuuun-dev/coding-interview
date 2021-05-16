package main

import (
	"fmt"
	"sort"
)

func calc(N int) {
	nums := toFloatAry(N)
	for _, v := range nums {
		if v == 1.1 {
			fmt.Printf("%d", 0)
		} else {
			fmt.Printf("%d", int(v))
		}
	}
	fmt.Println("")
}

func toFloatAry(N int) []float64 {
	var tmp []float64
	for N > 0 {
		var n float64
		n = float64(N % 10)
		if n == 0 {
			n = 1.1
		}
		tmp = append(tmp, n)
		N /= 10
	}
	sort.Float64s(tmp)
	return tmp
}

func main() {
	var N int
	fmt.Scan(&N)
	calc(N)
}
