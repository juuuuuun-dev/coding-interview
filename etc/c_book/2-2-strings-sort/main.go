package main

/*
n 文字の大文字文字列を先頭、または後方から一文字づつ辞書型比較して並び替える
同じだった場合はその次の文字を比較
*/
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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func calc(n int, str []string) {
	a, b := 0, n-1
	tmp := ""
	for a <= b {
		left := false
		// 先頭と後方を比較し、同じだった場合貧欲的に次の文字を捜査
		// 最後まで一緒だったら最後の文字を取得して次へ
		for i := 0; a+i <= b; i++ {
			fmt.Println(str[a+i], str[b-i])
			if str[a+i] < str[b-i] {
				left = true
				break
			} else if str[a+i] > str[b-i] {
				left = false
				break
			}
		}
		if left {
			tmp += str[a]
			a++
		} else {
			tmp += str[b]
			b--
		}
		fmt.Println("a", a, "b", b)
	}
	fmt.Println(tmp)
}

func main() {
	n := strToInt(readLine())
	str := strings.Split(readLine(), "")
	calc(n, str)
}
