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

func splitLine(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
}

func split(s string) []string {
	return strings.Fields(s)
}

func filterdData(lines [][]string, day int) []int {
	var start, end, max, min int
	for i := 0; i < len(lines); i++ {
		for n := 0; n < len(lines[i]); n++ {
			num := strToInt(lines[i][n])
			if i == 0 && n == 0 {
				start = num
			}
			if i == (day-1) && n == 1 {
				end = num
			}
			if n == 2 && num > max {
				max = num
			}
			if n == 3 && num < min || n == 3 && min == 0 {
				min = num
			}
		}

	}
	return []int{start, end, max, min}
}

func intAryToStrAry(nums []int) []string {
	var ret = make([]string, 0, len(nums))
	for _, num := range nums {
		var strVal = strconv.Itoa(num)
		ret = append(ret, strVal)
	}
	return ret
}

func main() {
	n := strToInt(readLine())
	lines := splitLine(n)
	filterd := filterdData(lines, n)
	fmt.Println(strings.Join(intAryToStrAry(filterd), " "))
}
