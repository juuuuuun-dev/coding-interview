package main

/*
1000 ~ 10000の数字で四則演算式を使って答えが回文になる数式を求める
*/
import (
	"fmt"
	"go/token"
	"go/types"
	"strconv"
)

func eval(str string) (types.TypeAndValue, error) {
	result, err := types.Eval(token.NewFileSet(), nil, token.NoPos, str)
	return result, err
}

func intToStr(i int) string {
	var strVal = strconv.Itoa(i)
	return strVal
}

func strToInt(s string) int {
	var intVal, e = strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return intVal
}

func reversInt(n int) int {
	ret := 0
	for n > 0 {
		remainder := n % 10
		ret *= 10
		ret += remainder
		n /= 10
	}
	return ret
}
func calc() {
	op := []string{"+", "-", "*", "/", ""}
	lenOp := len(op)
	for i := 1000; i < 10000; i++ {
		c := intToStr(i)
		for j := 0; j < lenOp; j++ {
			for k := 0; k < lenOp; k++ {
				for l := 0; l < lenOp; l++ {
					// 回文になるように後ろから追加
					val := string(c[3]) + op[j] + string(c[2]) + op[k] + string(c[1]) + op[l] + string(c[0])
					// 演算子が一つ以上
					if len(val) > 4 {
						ret, err := eval(val)
						if err == nil {
							if i == strToInt(ret.Value.String()) {
								fmt.Println(val)
								a := reversInt(i)
								fmt.Println(ret.Value, a, i)
							}
						}
					}
				}
			}
		}
	}

}

func calc2() {
	op := []string{"+", "-", "*", "/", ""}
	lenOp := len(op)
	for i := 100; i < 1000; i++ {
		c := intToStr(i)
		for j := 0; j < lenOp; j++ {
			for k := 0; k < lenOp; k++ {
				val := string(c[2]) + op[j] + string(c[1]) + op[k] + string(c[0])
				// fmt.Println(val)
				if len(val) > 3 {
					ret, err := eval(val)
					if err == nil {
						if i == strToInt(ret.Value.String()) {
							fmt.Println(val)
							a := reversInt(i)
							fmt.Println(ret.Value, a)
						}
					}
				}
			}
		}
	}

}
func main() {
	calc()
	// calc2()
	const aa = 1
	if aa == 1 {

	}
	// fmt.Println(eval("4*6/0+1"))
}
