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
func splitLine(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
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

func filterdData(n string, t int) []string {
	var ret []string
	for i := 0; i < t; i++ {
		num := readLine()
		if strings.Index(num, n) == -1 {
			ret = append(ret, num)
		}
	}
	return ret
}

func printData(data []string) {
	if len(data) > 0 {
		fmt.Println(strings.Join(data, "\n"))
	} else {
		fmt.Println("none")
	}
}
func main() {
	n := readLine()
	t := strToInt(readLine())
	data := filterdData(n, t)
	printData(data)
}
