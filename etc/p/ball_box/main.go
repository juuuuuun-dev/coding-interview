package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func lineSplit(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
}

func filterdBox(lines [][]string, r int) []int {
	var slice []int
	for i := 0; i < len(lines); i++ {
		isIn := true
		for _, n := range lines[i] {
			if strToInt(n) < r*2 {
				isIn = false
			}
		}
		if isIn {
			slice = append(slice, i+1)
		}
	}
	return slice
}

func main() {
	s := readLine()
	firstLine := split(s)
	var b, r int = strToInt(firstLine[0]), strToInt(firstLine[1])
	lines := lineSplit(b)
	filterd := filterdBox(lines, r)
	for _, val := range filterd {
		fmt.Println(val)
	}
}
