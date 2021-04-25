package main

import (
	"bufio"
	"fmt"
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

func s2i(s string) int {
	var i, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := s2i(readLine())
	total := calc(n)
	fmt.Println(total)
}

func calc(n int) int {
	var t, ab int
	var slice []int
	for i := 0; i < n; i++ {
		o := s2i(readLine())
		slice = append(slice, o)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			ab = abs(1 - slice[i])
		} else {
			ab = abs(slice[i] - slice[i-1])
		}
		t = t + ab
	}
	return t
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
