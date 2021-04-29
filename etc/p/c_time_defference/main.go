package main

import (
	"bufio"
	"fmt"
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

func splitLine(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
}

func intToStr(i int) string {
	var strVal = strconv.Itoa(i)
	return strVal
}

func withZero(n int) string {
	if n < 10 {
		return "0" + intToStr(n)
	}
	return intToStr(n)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func calc(lines [][]string, c []string) {
	var target int
	times := strings.Split(c[1], ":")
	hour := strToInt(times[0])
	subHour := hour
	if hour == 0 {
		subHour = 24
	}
	minute := strToInt(times[1])
	for _, v := range lines {
		for i := 0; i < 2; i++ {
			if v[0] == c[0] {
				target = strToInt(v[1])
			}
		}
	}

	for _, v := range lines {
		var sub int
		for i := 0; i < 2; i++ {
			sub = strToInt(v[1]) - target
		}
		stTime := sub + subHour
		// fmt.Println("stTime", stTime)
		if stTime >= 24 {
			stTime = stTime - 24
		}
		if stTime < 0 {
			stTime = 24 - abs(stTime)
		}
		fmt.Println(withZero(stTime) + ":" + withZero(minute))
	}
}

func main() {
	n := strToInt(readLine())
	lines := splitLine(n)
	c := split(readLine())
	calc(lines, c)
}
