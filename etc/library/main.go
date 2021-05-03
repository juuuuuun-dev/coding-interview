package main

import (
	"bufio"
	"go/token"
	"go/types"
	"math"
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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func splitLine(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
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

func intToStr(i int) string {
	var strVal = strconv.Itoa(i)
	return strVal
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

func intAryToStrAry(nums []int) []string {
	var ret = make([]string, 0, len(nums))
	for _, num := range nums {
		var strVal = strconv.Itoa(num)
		ret = append(ret, strVal)
	}
	return ret
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func floor(a int) int {
	return int(math.Floor(float64(a)))
}

func withZero(n int) string {
	if n < 10 {
		return "0" + intToStr(n)
	}
	return intToStr(n)
}

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

func eval(str string) (types.TypeAndValue, error) {
	result, err := types.Eval(token.NewFileSet(), nil, token.NoPos, str)
	return result, err
}

// 文字列をindexで表示
// fmt.Println(string("Hello"[1]))

// sort
// sort.Ints(keys)
// sort.Sort(sort.IntSlice(nums))
// sort.Sort(sort.Reverse(sort.IntSlice(nums)))
