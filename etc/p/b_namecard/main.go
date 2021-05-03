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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
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

func calc(n []int) {
	pockets := n[0] * 2
	number := n[1] - 1
	page := number / pockets
	position_sheet := number % pockets
	revers_position := pockets - position_sheet
	position := page*pockets + revers_position
	fmt.Println(position)
}

func main() {
	n := strAryToIntAry(split(readLine()))
	calc(n)
}
