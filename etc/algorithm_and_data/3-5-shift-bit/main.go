package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func split(s string) []string {
	return strings.Fields(s)
}

func strAryToIntAry(strs []string) []int {
	var ret = make([]int, 0, len(strs))
	for _, str := range strs {
		var intVal, e = strconv.Atoi(string(str))
		if e != nil {
			panic(e)
		}
		ret = append(ret, intVal)
	}
	return ret
}

func calc(n int, m int, l []int) int {
	a := 6
	// var tmp []int
	// flag := false
	for bit := 0; bit < (1 << n); bit++ {
		bit += 1
		fmt.Println("bit", bit)
		sum := 0
		for i := 0; i < n; i++ {
			if bit>>i&1 == 1 {
				fmt.Println("i", i)
				sum += l[i]
			}
		}
		// if sum == m {
		// 	flag = true
		// }
		fmt.Println("sum", sum)
	}
	fmt.Printf("aa %b ", 1<<3)
	return a
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	// l := strAryToIntAry(split(readLine()))
	// a := calc(n, m, l)
	// fmt.Println(a)

}
