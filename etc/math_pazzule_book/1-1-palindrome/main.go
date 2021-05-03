package main

import (
	"fmt"
	"strconv"
)

// 10,8,2進数で10進数で10以上の数字で回文となる数字を求める
// 2進数で一桁が1は奇数。0は偶数。なので奇数

// revers int
func reversInt(n int) int {
	ret := 0
	for n > 0 {
		remainder := n % 10
		ret *= 10
		ret += remainder
		n /= 10
	}
	return ret
}

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func calc() {
	n := 11
	for true {
		s2 := strToInt(fmt.Sprintf("%b", n))
		s8 := strToInt(fmt.Sprintf("%o", n))
		if n == reversInt(n) && s2 == reversInt(s2) && s8 == reversInt(s8) {
			fmt.Println(n, s2, s8)
			break
		}
		n += 2
	}
}

func main() {
	calc()
}
