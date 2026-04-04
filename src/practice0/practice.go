package practice0

import (
	"fmt"
)

/*
課題0

2つの整数を入力させ、合計を表示する。
*/
func Run() {
	fmt.Print("1つめの数は？")
	var n1 int
	fmt.Scan(&n1)

	fmt.Print("2つめの数は？")
	var n2 int
	fmt.Scan(&n2)

	sum := n1 + n2

	fmt.Printf("合計：%d\n", sum)
}
