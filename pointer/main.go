package main

import "fmt"

// go 中与指针相关的两个操作符： 取地址符&，取地址的值*
// 指针变量的值是一个内存地址，指针变量本身有自己的内存地址。
// 取地址符&，获取一个变量的地址。
// 取地址的值*，获取指向地址的值。

func main() {
	num := 3
	fmt.Println(&num)
	fmt.Println(*&num)       // 只是示范作用，取地址的值
	fmt.Printf("%T\n", &num) // 查看地址类型 *int，即int类型的指针

	var p *int     // nil pointer
	fmt.Println(p) // <nil>

	p2 := new(int)   // 指向一个新分配的内存地址
	fmt.Println(p2)  // 打印p2指向的内存地址
	fmt.Println(*p2) // 零值 0

	*p2 = 10         // 赋值
	fmt.Println(*p2) // 10
}
