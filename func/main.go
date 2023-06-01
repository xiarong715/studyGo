package main

import "fmt"

func testDefer() {
	defer func() { fmt.Println("111") }() // 先注册的defer函数后执行，可想像成注册的defer函数放入了一个栈中。先进后出。
	defer func() { fmt.Println("222") }()
	fmt.Println("testDefer")
}

func testVariableParam(str string, y ...int) {
	fmt.Println(str, y)
}

func add(x, y int) int {
	return x + y
}

// 函数作为函数的参数
func testFuncParam(a, b int, addfunc func(x, y int) int) int {
	return addfunc(a, b) // 只有函数，没有引用外部环境，不是一个闭包
}

// 函数作为返回值
func testFuncReturnValue() func(int, int) int {
	return add
}

// 匿名函数，多用于实现回调函数和闭包
// 在函数体内，只能定义为匿名函数
func testAnonimityFunc() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2)) //通过函数变量调用函数

	res := func(x, y int) int {
		return x - y
	}(1, 2) // 定义匿名函数后，马上调用
	fmt.Println(res)
}

// 闭包，闭包是一个函数，这个函数包含了他外部作用域的一个变量。
// 闭包 = 函数 + 引用环境
func adder(x int) func(int) int {
	return func(y int) int { // 匿名函数  和 外部变量 x 组成一个闭包
		x += y
		return x
	}
}

func main() {
	testDefer()
	testVariableParam("hello") // 可变参数可不填
	testVariableParam("world", 1, 2, 3)
	res := testFuncParam(1, 2, add)
	fmt.Println(res)
	res = testFuncReturnValue()(1, 5) // 返回函数后直接调用
	fmt.Println(res)
	testAnonimityFunc()
	res = adder(2)(3) // adder函数返回一个闭包，继续调用闭包
	fmt.Println(res)
}
