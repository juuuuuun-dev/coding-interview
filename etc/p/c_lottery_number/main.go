package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

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

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func split(s string) []string {
	return strings.Fields(s)
}

func sToI(s string) int {
	var i, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func prepare() ([]int, int) {
	first := readLine()
	winning := strAryToIntAry(split(first))
	cardNum := sToI(readLine())
	return winning, cardNum
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

func calc(winning []int, cardNum int, lines [][]string) {
	for _, v := range lines {
		line := strAryToIntAry(v)
		count := 0
		for _, n := range line {
			for _, w := range winning {
				if n == w {
					count += 1
				}
			}
		}
		fmt.Println(count)
	}
}

func main() {
	winning, cardNum := prepare()
	lines := splitLine(cardNum)
	calc(winning, cardNum, lines)
}
