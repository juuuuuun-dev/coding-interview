package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sortReverse(n []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
}

func calc(N, TOTAL int, A []int) {
	sortReverse(A)
	fmt.Println(A)
	used := map[int]bool{}
	combi := [][]int{}
	t := TOTAL / 2
	fmt.Println("TOTAL", TOTAL)
	fmt.Println("t", t)
	for k, v := range A {
		tmp := []int{}
		if _, ok := used[k]; ok == false {
			tmp = append(tmp, v)
			tmpT := v
			used[k] = true
			for i := 1; i <= N-1; i++ {
				if _, ok := used[i]; ok == false {
					if tmpT+A[i] < t && len(used) < len(A)-1 {
						tmpT += A[i]
						fmt.Println("tmpT", tmpT)
						tmp = append(tmp, A[i])
						used[i] = true
					}
				}
			}
		}
		if len(tmp) > 0 {
			combi = append(combi, tmp)
		}
	}
	if len(combi) <= 3 {
		fmt.Println(combi)
		fmt.Println("YES")
		for _, v := range combi {
			for _, n := range v {
				fmt.Printf("%d ", n)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var r = bufio.NewReader(os.Stdin)
	var N, TOTAL int
	fmt.Fscan(r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &A[i])
		TOTAL += A[i]
	}
	calc(N, TOTAL, A)
}
