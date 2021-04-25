package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type strType map[string]string

func readLine() (s string) {
	if sc.Scan() {
		s = sc.Text()
	}
	return s
}

func main() {
	s := readLine()
	strs := strs()
	res := replace(s, strs)
	fmt.Println(res)
}

func strs() strType {
	m := strType{"a": "a", "i": "i", "u": "u", "e": "e", "o": "o", "A": "A", "I": "I", "U": "U", "E": "E", "O": "O"}
	return m
}

func replace(s string, strs strType) (r string) {
	slice := strings.Split(s, "")
	len := len(slice)
	for i := 0; i < len; i++ {
		if _, has := strs[slice[i]]; has == false {
			r = r + slice[i]
		}
	}
	return r
}
