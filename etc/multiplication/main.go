package main

import "fmt"

// 2つの数の掛け算からなる数字nの内訳
func multi(n int) [][]int {
	res := [][]int{}
	for i := 1; i*i <= n; i++ {
		// iで割り切れるか
		if n%i != 0 {
			continue
		}
		tmp := []int{}
		tmp = append(tmp, i)
		tmp = append(tmp, n/i)
		res = append(res, tmp)
	}
	return res
}

func main() {
	fmt.Println(multi(10000))
}
