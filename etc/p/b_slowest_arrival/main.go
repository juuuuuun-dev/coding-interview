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

func floor(a int) int {
	return int(math.Floor(float64(a)))
}

func withZero(n int) string {
	if n < 10 {
		return "0" + intToStr(n)
	}
	return intToStr(n)
}

func filterdMinute() int {
	times := strAryToIntAry(split(readLine()))
	n := strToInt(readLine())
	lines := splitLine(n)
	limitMinute := 60 * 9 // 9時を分に変換
	var minute, pass int

	for i := 0; i < n; i++ {
		hourTimeTable := strToInt(lines[i][0])
		minutesTimeTable := strToInt(lines[i][1])
		toStation := hourTimeTable*60 + minutesTimeTable
		minute = toStation
		minute += times[1]
		minute += times[2]
		if minute < limitMinute {
			pass = toStation - times[0]
		} else {
			break
		}
	}
	return pass
}

func printTime(n int) {
	hour := withZero((n / 60))
	minute := withZero(n % 60)
	fmt.Println(hour + ":" + minute)
}

func main() {
	pass := filterdMinute()
	printTime(pass)
}
