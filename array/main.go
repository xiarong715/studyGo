package main

import "fmt"

// 数组的长度是数组类型的一部分。
// 定义数组时必须指定存放元素的类型和长度。

func main() {
	// 定义
	var a1 [3]int
	var a2 [4]int
	fmt.Printf("%T\n", a1) // 数组a1的类型：[3]int
	fmt.Printf("%T\n", a2) // 数组a2的类型：[4]int

	// 初始化
	// 定义数组时，元素值默认为零值：整型和浮点型为0，bool类型为false，字符串类型为""
	fmt.Println(a1) // [0 0 0]
	fmt.Println(a2) // [0 0 0 0]

	// 初始化方式1：在大括号中放入与长度一样多的元素
	a1 = [3]int{1, 2, 3}
	fmt.Println(a1) // [1 2 3]

	// 初始化方式2：根据初始值依次初始化各个元素，不够则补上默认零值
	a2 = [4]int{2, 3, 4}
	fmt.Println(a2) // [2 3 4 0]

	// 初始化方式3：根据初始值个数自动识别数组长度 [...]
	a3 := [...]int{2, 3, 4}
	fmt.Printf("%T %v\n", a3, a3) // [3]int  [2 3 4]

	// 初始化方式4：根据索引初始化
	a4 := [5]int{0: 1, 2: 3, 4: 5}
	fmt.Println(a4)

	// 取值
	// for i
	for i := 0; i < len(a4); i++ {
		fmt.Print(a4[i], " ")
	}
	fmt.Println()

	// for range
	for _, v := range a3 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// 多维数组：可理解为元素为数组的数组
	// 定义多维数组
	var darry = [2][3]int{
		{1, 2, 3},
		{2, 3, 4},
	}
	fmt.Println(darry)

	// 多维数组编历
	// 双for range
	for _, v1 := range darry {
		for _, v := range v1 {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	// 数组是值类型，不是引用类型
	a5 := [3]int{0, 1, 2}
	// 值拷贝，两数组没关联。有自己独立的空间。。。。切片的类型是一个结构体，有一个字段指向同一个底层数组，所以切片相互赋值时，两个切片会指向同一个底层数组。
	// 但数组是值类型，赋值时，直接拷贝了值，两数组没联系。
	// go 都是值传递。
	a6 := a5
	a6[0] = 100
	fmt.Println(a5, a6) // [0 1 2] [100 1 2]

	// 数组实践
	// 打印出两个元素和为8的下标
	a7 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a7); i++ {
		for j := i + 1; j < len(a7); j++ {
			if a7[i]+a7[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
