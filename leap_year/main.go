package main
import "fmt"

func main() {
	years := []int{4, 1000, 1992, 2000, 2001}
	for _, v := range years {
		if v == 4 {
			continue
		}
		if v % 400 == 0 || (v % 4 == 0 && v % 100 != 0) {
			str := fmt.Sprintf("%d is a leap year", v)
			fmt.Println(str)
		} else {
			str := fmt.Sprintf("%d is not a leap year", v)
			fmt.Println(str)
		}
	}
}