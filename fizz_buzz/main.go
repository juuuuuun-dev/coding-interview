package main

import "fmt"

func main() {
	count := 100
	for i := 1; i < count; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("Fizz Buzz")
		}
		if i % 3 == 0 {
			fmt.Println("Fizz")
		}
		if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}