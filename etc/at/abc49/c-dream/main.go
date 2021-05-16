package main

import "fmt"

func rec(S string, strs [4]string) bool {
	l := len(S)
	if l == 0 {
		return true
	}
	for _, v := range strs {
		vl := len(v)
		if l >= vl {
			t := S[l-vl : l]
			if t == v {
				S = S[0 : l-vl]
				return rec(S, strs)
			}
		}
	}
	return false
}

func main() {
	var S string
	fmt.Scanf("%s", &S)
	strs := [4]string{"dream", "dreamer", "erase", "eraser"}
	res := rec(S, strs)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
