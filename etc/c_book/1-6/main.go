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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func max(s []int) int {
	var tmp int
	len := len(s)
	tmp = s[0]
	for i := 1; i < len; i++ {
		if s[i] > tmp {
			tmp = s[i]
		}
	}
	return tmp
}

func calc(n int, nums []int) int {
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				len := nums[i] + nums[j] + nums[k]
				var maxTmp []int
				maxTmp = append(maxTmp, nums[i])
				maxTmp = append(maxTmp, nums[j])
				maxTmp = append(maxTmp, nums[k])
				max := max(maxTmp)
				rest := len - max
				if rest > max && len > ans {
					ans = len
					fmt.Print(maxTmp)
				}
			}
		}
	}
	return ans
}

func main() {
	n := strToInt(readLine())
	nums := strAryToIntAry(split(readLine()))
	ans := calc(n, nums)
	fmt.Println(ans)
}
