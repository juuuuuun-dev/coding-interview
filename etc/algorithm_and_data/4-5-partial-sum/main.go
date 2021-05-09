package main

/*
部分和問題
N個の整数を足してWを作れるか？

N W
N1 N2 N3 N4 ...Nn

考え方
WからNnを引いてって0になるか

*/
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

func calc(n int, w int, a []int) bool {
	if n == 0 {
		if w == 0 {
			return true
		}
		return false
	}

	// 引かないパターン
	if calc(n-1, w, a) {
		return true
	}
	// a[n-1]を引くパターン
	if calc(n-1, w-a[n-1], a) {
		return true
	}
	return false
}

func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)
	lines := strAryToIntAry(split(readLine()))
	res := calc(N, W, lines)
	fmt.Println(res)
}
