package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func main() {
	s := readLine()
	strLen := len(s)
	decorationLen := 2 + strLen
	decorationStr := "+"
	decoration := makeDecoration(decorationLen, decorationStr)
	fmt.Println(decoration)
	fmt.Println(decorationStr + s + decorationStr)
	fmt.Println(decoration)
}

func makeDecoration(len int, str string) (s string) {
	for i := 0; i < len; i++ {
		s += str
	}
	return s
}
