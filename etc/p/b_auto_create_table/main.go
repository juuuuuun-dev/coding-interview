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

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func prepare() (h int, w int) {
	n := split(readLine())
	h = strToInt(n[0])
	w = strToInt(n[1])
	return h, w
}

func intAryToStrAry(nums []int) []string {
	var ret = make([]string, 0, len(nums))
	for _, num := range nums {
		var strVal = strconv.Itoa(num)
		ret = append(ret, strVal)
	}
	return ret
}

// 縦(col)を作成してから横(row)
func createTable(h int, w int, lines [][]string) {
	col1Start := strToInt(lines[0][0])
	col1Sub := strToInt(lines[1][0]) - col1Start
	col1Slice := createCell(col1Start, col1Sub, h)

	col2Start := strToInt(lines[0][1])
	col2Sub := strToInt(lines[1][1]) - col2Start
	col2Slice := createCell(col2Start, col2Sub, h)

	for i := 0; i < h; i++ {
		rowStart := col1Slice[i]
		rowSub := col2Slice[i] - rowStart
		rowSlice := createCell(rowStart, rowSub, w)
		fmt.Println(strings.Join(intAryToStrAry(rowSlice), " "))
	}
}

// セル計算式 start + i * sub
func createCell(start int, sub int, n int) []int {
	var ret []int
	for i := 0; i < n; i++ {
		// fmt.Println(start, "+", i, "*", sub, "=", start+i*sub)
		ret = append(ret, start+i*sub)
	}
	return ret
}

func main() {
	h, w := prepare()
	lines := splitLine(2)
	createTable(h, w, lines)
}
