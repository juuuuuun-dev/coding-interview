package main

import "fmt"

func main() {
	var n int = 100
	var p *int         //intのポインタ型
	type hogeType int  // 型宣言
	var hoge *hogeType // hogeTypeのポインタ型
	p = &n             // ポインタ型の変数
	fmt.Println(n)
	fmt.Println(hoge)
	fmt.Println(&hoge) // memory address
	fmt.Println(&n)    // memory address
	fmt.Println(p)     // &nと同じ値が表示されます
	fmt.Println(*p)    // ポインタ型の値
	*p = 50            // デリファレンス
	fmt.Println(n)     // 50
	one(p)
	fmt.Println(n)
}

func one(x *int) {
	*x = 1
}
