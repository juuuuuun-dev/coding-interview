package main

import "fmt"

func binarySearch(n int, slice []int) (int, int) {
	l, r := 0, len(slice)
	count := 0
	for r-1 >= l {
		count++
		mid := (l + r) / 2
		if n == slice[mid] {
			return n, count
		} else if n > slice[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1, count
}

func main() {
	var n, s, e int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d %d", &s, &e)
	slice := make([]int, e-s)
	for i := 0; i < e-s; i++ {
		slice[i] = i + s
	}
	a, c := binarySearch(n, slice)
	fmt.Println(a, c)
}
