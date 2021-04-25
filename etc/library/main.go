package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func sToI(s string) int {
	var i, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
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
