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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func calc(n int, l []int) {
	ans := 0
	fmt.Println(l)
	for n > 1 {
		min1, min2 := 0, 1
		fmt.Println("l[min1]", l[min1], "l[min2]", l[min2])
		if l[min1] > l[min2] {
			min1, min2 = min2, min1
		}

		fmt.Println("min1", min1, "min2", min2)
		// 3つめ以降と比較し最小を探す
		for i := 2; i < n; i++ {
			fmt.Println("i", i)
			if l[i] < l[min1] {
				min2 = min1
				min1 = i
			} else if l[i] < l[min2] {
				min2 = i
			}
		}
		t := l[min1] + l[min2]
		fmt.Println("t", t)
		ans += t
		if min1 == n-1 {
			min1, min2 = min2, min1
		}
		l[min1] = t
		l[min2] = l[n-1]
		n--
	}
	fmt.Println(ans)
}

func main() {
	n := strToInt(readLine())
	l := strAryToIntAry(split(readLine()))
	calc(n, l)
}
