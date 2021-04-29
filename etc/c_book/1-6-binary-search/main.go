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

func binary_search(n int, x int, nums []int) bool {
	l := 0
	r := n
	for r-1 >= l {
		mid := (l + r) / 2
		if nums[mid] == x {
			return true
		} else if nums[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}

func calc(n int, x int, nums []int) bool {
	sort.Sort(sort.IntSlice(nums))
	has := false
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				// x - nums[a] - nums[b] - nums[b] の値がnumsに存在するか
				if binary_search(n, x-nums[a]-nums[b]-nums[c], nums) {
					has = true
				}
			}
		}
	}
	return has
}

func main() {
	n := strToInt(readLine())
	x := strToInt(readLine())
	nums := strAryToIntAry(split(readLine()))
	has := calc(n, x, nums)
	fmt.Println(has)
}
