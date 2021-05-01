package main

/*
a の金額を最小のコインの組み合わせで支払う
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

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
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

func min(s []int) int {
	var tmp int
	len := len(s)
	tmp = s[0]
	for i := 1; i < len; i++ {
		if s[i] < tmp {
			tmp = s[i]
		}
	}
	return tmp
}

func calc(a int, n []int, conins [6]int) {
	ans := 0
	for i := 5; i >= 0; i-- {
		var tmp []int
		tmp = append(tmp, n[i])
		tmp = append(tmp, a/conins[i])
		minC := min(tmp)
		a -= conins[i] * minC
		ans += minC
	}
	fmt.Println(ans)
}

func main() {
	a := strToInt(readLine())
	n := strAryToIntAry(split(readLine()))
	var coins [6]int = [6]int{1, 5, 10, 50, 100, 500}
	calc(a, n, coins)
}
