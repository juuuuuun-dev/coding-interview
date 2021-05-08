package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func dfs(i int, str string, len int, nums []string) int {
	if i == len-1 {
		n := strings.Split(str, "+")
		var t int
		for _, v := range n {
			t += strToInt(v)
		}
		return t
	}
	// + を追加パターンとそのままパターン
	return dfs(i+1, str+nums[i+1], len, nums) + dfs(i+1, str+"+"+nums[i+1], len, nums)
}

func main() {
	var N string
	fmt.Scanf("%s", &N)
	nums := strings.Split(N, "")
	len := len(nums)
	res := dfs(0, nums[0], len, nums)
	fmt.Println(res)
}
