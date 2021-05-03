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

func calc(n string) {
	xIndex := strings.Index(n, "x")
	plusIndex := strings.Index(n, "+")
	strs := strings.Split(n, "")
	// n + n = x
	if xIndex == 4 && plusIndex == 1 {
		fmt.Println(strToInt(strs[0]) + strToInt(strs[2]))
	}
	// n - n = x
	if xIndex == 4 && plusIndex == -1 {
		fmt.Println(strToInt(strs[0]) - strToInt(strs[2]))
	}
	// x + n = n
	if xIndex == 0 && plusIndex == 1 {
		fmt.Println(strToInt(strs[4]) - strToInt(strs[2]))
	}
	// x - n = n
	if xIndex == 0 && plusIndex == -1 {
		fmt.Println(strToInt(strs[4]) + strToInt(strs[2]))
	}
	// n + x = n
	if xIndex == 2 && plusIndex == 1 {
		fmt.Println(strToInt(strs[4]) - strToInt(strs[0]))
	}
	// n - x = n
	if xIndex == 2 && plusIndex == -1 {
		fmt.Println(strToInt(strs[0]) - strToInt(strs[4]))
	}
}

func main() {
	n := strings.Replace(readLine(), " ", "", -1)
	calc(n)
}
