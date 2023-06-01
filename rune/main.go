package main

import "fmt"

func main() {
	r := "夏"      // 一个汉字3个字节
	r2 := "hello" // 长度为5
	fmt.Println(len(r))
	fmt.Println(len(r2))

	r3 := "夏hello"
	for i := 0; i < len(r3); i++ {
		fmt.Printf("%c", r3[i]) // 有乱码
	}
	fmt.Println()

	for _, c := range r3 {
		fmt.Printf("%c", c) // 可以正常输出中文
	}
	fmt.Println()

	// 修改字符串
	r4 := "小白兔"
	r5 := []rune(r4)        // 字符串不能直接修改，需转换成rune切片再修改
	r5[0] = '大'             // 修改rune切片的值
	fmt.Println(string(r5)) // 大白兔        成功修改了字符串

	// rune是int32的别名
	// 双引号是string类型
	// 单引号是int32类型。rune和中、英文没关系，只要单引号，都是int32类型
	r6 := "夏"
	r7 := '夏'
	r8 := '3'
	fmt.Printf("%T %T %T\n", r6, r7, r8) // string int32 int32
}
