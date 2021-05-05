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

func split(s string) []string {
	return strings.Fields(s)
}

func calc(n int, m int, a []int, b []int) {
	min := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := a[i] + b[j]
			if tmp < m {
				continue
			} else {
				if tmp < min || min == 0 {
					min = tmp
				}
			}
		}
	}
	fmt.Println(min)
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	a := strAryToIntAry(split(readLine()))
	b := strAryToIntAry(split(readLine()))
	calc(n, m, a, b)
}
