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

type kv struct {
	Key   int
	Value int
}

func calc(n int, k int, nums []int) {
	m := map[int]int{}
	count := 0
	for _, val := range nums {
		m[val] += 1
	}
	len := len(m)
	if len > k {
		var ss []kv
		for k, v := range m {
			ss = append(ss, kv{k, v})
		}
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value < ss[j].Value
		})
		for _, value := range ss {
			if k < len {
				for i := 0; i < value.Value; i++ {
					len -= 1
					count += 1
				}
			}
		}
	}
	fmt.Println("count", count)
}

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)
	nums := strAryToIntAry(split(readLine()))
	calc(N, K, nums)
}
