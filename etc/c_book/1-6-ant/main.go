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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
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

func max(s []int) int {
	var tmp int
	len := len(s)
	tmp = s[0]
	for i := 1; i < len; i++ {
		if s[i] > tmp {
			tmp = s[i]
		}
	}
	return tmp
}

func min(s []int) int {
	var tmp int
	len := len(s)
	tmp = s[0]
	for i := 1; i < len; i++ {
		if s[i] < tmp {
			tmp = s[i]
		}
	}
	return tmp
}

func calcMin(L int, n int, ants []int) int {
	var tmp, ret = 0, 0
	for i := 0; i < n; i++ {
		var tmpMin []int
		tmpMin = append(tmpMin, L-ants[i])
		tmpMin = append(tmpMin, ants[i])
		tmp = min(tmpMin) // 左右向きどちらから近い方を取得
		if tmp > ret {
			ret = tmp //最短の中からもっとも距離のあるもの
		}
	}
	return ret
}

func calcMax(L int, n int, ants []int) int {
	var tmp, ret = 0, 0
	for i := 0; i < n; i++ {
		var tmpMin []int
		tmpMin = append(tmpMin, L-ants[i])
		tmpMin = append(tmpMin, ants[i])
		tmp = max(tmpMin) // 左右向きどちらから遠い方を取得
		if tmp > ret {
			ret = tmp //最長の中からもっとも距離のあるもの
		}
	}
	return ret
}

func main() {
	L := strToInt(readLine())
	n := strToInt(readLine())
	ants := strAryToIntAry(split(readLine()))
	min := calcMin(L, n, ants)
	max := calcMax(L, n, ants)
	fmt.Println(min, max)
}
