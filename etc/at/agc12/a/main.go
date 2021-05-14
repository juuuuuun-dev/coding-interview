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
	i := 2
	ans := []int{}
	for len(ans) < N {
		ans = append(ans, nums[i-1])
		i += 2
	}
	total := 0
	for _, v := range ans {
		total += v
	}
	fmt.Println(total)
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	nums := strAryToIntAry(split(readLine()))
	calc(N, nums)
}
