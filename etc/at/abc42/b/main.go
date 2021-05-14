package main

import (
	"fmt"
	"sort"
)

func calc(strs []string) {
	sort.Strings(strs)
	for _, v := range strs {
		fmt.Printf("%s", v)
	}
}

func main() {
	var N, L int
	fmt.Scanf("%d %d", &N, &L)
	strs := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &strs[i])
	}
	calc(strs)
}
