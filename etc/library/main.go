package main

import (
	"bufio"
	"go/token"
	"go/types"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

var r = bufio.NewReader(os.Stdin)

// 一行 space int
// A := make([]int, N)
// for i := 0; i < N; i++ {
// 	fmt.Fscan(r, &A[i])
// }

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

func splitLineToInt(n int) [][]int {
	var slice [][]int
	for i := 0; i < n; i++ {
		r := readLine()
		l := strAryToIntAry(split(r))
		slice = append(slice, l)
	}
	return slice
}
func splitLineToStr(n int) [][]string {
	var slice [][]string
	for i := 0; i < n; i++ {
		r := readLine()
		l := split(r)
		slice = append(slice, l)
	}
	return slice
}

func maxSlice(s []int) int {
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

func max(n, i int) int {
	return int(math.Max(float64(n), float64(i)))
}

func min(n, i int) int {
	return int(math.Min(float64(n), float64(i)))
}

func minSlice(s []int) int {
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

func intToStr(i int) string {
	var strVal = strconv.Itoa(i)
	return strVal
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

func intAryToStrAry(nums []int) []string {
	var ret = make([]string, 0, len(nums))
	for _, num := range nums {
		var strVal = strconv.Itoa(num)
		ret = append(ret, strVal)
	}
	return ret
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func floor(a int) int {
	return int(math.Floor(float64(a)))
}

// 数字を配列に
func intToIntoAry(N int) []int {
	var res []int
	for N > 0 {
		res = append(res, N%10)
		N /= 10
	}
	return res
}

func withZero(n int) string {
	if n < 10 {
		return "0" + intToStr(n)
	}
	return intToStr(n)
}

// revers int 数字反転
// 先頭が0になる場合は追加しない
// 10で割った余りで一桁台の数字を取得
// その余りを答えに追加して10を掛ける
// 対象の数字を10で割る。
// これを繰り返す
// 対象の数字を10で割るを繰り返すことによって数字の桁分処理される
// 桁の数だけ答えに10を掛けるので最終的に同じ桁になる
func reversInt(n int) int {
	res := 0
	for n > 0 {
		remainder := n % 10
		res *= 10
		res += remainder
		n /= 10
	}
	return res
}

func eval(str string) (types.TypeAndValue, error) {
	result, err := types.Eval(token.NewFileSet(), nil, token.NoPos, str)
	return result, err
}

// 文字列をindexで表示
// fmt.Println(string("Hello"[1]))

// sort
// sort.Ints(keys)
// sort.Sort(sort.IntSlice(nums))
// sort.Sort(sort.Reverse(sort.IntSlice(nums)))

func sortReverse(n []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
}

// 深さ優先 組み合わせ 再帰的

// func dfs(i int, str string, len int, nums []string) int {
// 	if i == len-1 {
// 		n := strings.Split(str, "+")
// 		var t int
// 		for _, v := range n {
// 			t += strToInt(v)
// 		}
// 		return t
// 	}
// 	// + を追加パターンとそのままパターン
// 	return dfs(i+1, str+nums[i+1], len, nums) + dfs(i+1, str+"+"+nums[i+1], len, nums)
// }

// 数字の各桁をたす
func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
