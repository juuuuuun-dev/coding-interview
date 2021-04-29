package main

import (
	"bufio"
	"fmt"
	"os"
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

func max(s []int) int {
	tmp := s[0]
	len := len(s)
	for i := 1; i < len; i++ {
		if s[i] > tmp {
			tmp = s[i]
		}
	}
	return tmp
}

func prepare() [][]string {
	n := split(readLine())
	var temp [][]string
	temp = append(temp, strings.Split(n[0]+n[1], ""))
	temp = append(temp, strings.Split(n[1]+n[0], ""))
	return temp
}

func convertName(names [][]string) int {
	aToZ := "abcdefghijklmnopqrstuvrsqz"
	var nums []int
	for _, v := range names {
		var base []int
		var index int
		for _, n := range v {
			index = strings.Index(aToZ, n) + 1
			base = append(base, index)
		}
		nums = append(nums, calc(base)[0])
	}
	fmt.Println(nums)
	return max(nums)
}

func calc(n []int) []int {
	len := len(n) - 1
	var tmp []int
	fmt.Println(n)
	if len > 0 {
		for i := 0; i < len; i++ {
			n := (n[i] + n[i+1]) % 101
			tmp = append(tmp, n)
		}
		return calc(tmp)
	} else {
		return n
	}
}

func main() {
	names := prepare()
	num := convertName(names)
	fmt.Printf("%v\n", num)
	fmt.Println(160 % 2)
}
