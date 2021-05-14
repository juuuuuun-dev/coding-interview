package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func sortReverse(n []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
}

func calc(N int, nums []int) {
	sortReverse(nums)
	a := 0
	b := 0
	for i := 0; i < N; i++ {
		fmt.Println("i%2", i%2)
		if i%2 == 0 {
			a += nums[i]
		} else {
			b += nums[i]
		}
	}
	fmt.Println(a - b)
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	nums := strAryToIntAry(split(readLine()))
	calc(N, nums)
}
