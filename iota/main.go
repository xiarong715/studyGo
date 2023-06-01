package main

import (
	"fmt"
)

func test1() {
	const (
		d1 = iota // 0
		d2        // 1
		d3        // 2
		d4        // 3
	)
	fmt.Println(d1, d2, d3, d4)
}

func test2() {
	const (
		d1 = iota // 0
		d2        // 1
		_         // 2
		d3        // 3
	)
	fmt.Println(d1, d2, d3)
}

func test3() {
	const (
		d1 = iota // 0
		d2 = 100  // 100
		d3        // 100
		d4        // 100
	)
	fmt.Println(d1, d2, d3, d4)
}

func test4() {
	const (
		d1 = iota // 0
		d2 = 100  // 100
		d3 = iota // 2
		d4        // 3
	)
	fmt.Println(d1, d2, d3, d4)
}

func test5() {
	const (
		d1, d2 = iota + 1, iota + 2 // 1 2
		d3, d4 = iota + 2, iota + 3 // 3 4
	)
	fmt.Println(d1, d2, d3, d4)
}

func test6() {
	const (
		_  = iota
		KB = 1 << (10 * iota) // 2^10 = 1024
		MB = 1 << (10 * iota) // 2^20 = 1048576
		GB = 1 << (10 * iota) // 2^30 = 1073741824
		TB = 1 << (10 * iota) // 2^40 = 1099511627776
		PB = 1 << (10 * iota) // 2^50 = 1125899906842624
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}

func test7() {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}

// iota 首次出现是值为0
// iota 之后每添加一行值+1
// const 变量没有表达式赋值时，表达式同上一行。

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	fmt.Println("hello iota")
}
