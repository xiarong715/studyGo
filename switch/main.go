package main

import "fmt"

// switch 比 if 更简洁，可读性更好
// case 后支持多个参数
// case 后加判断
func main() {
	finger := 2
	if finger == 1 {
		fmt.Println("小姆指")
	} else if finger == 2 {
		fmt.Println("无名指")
	} else if finger == 3 {
		fmt.Println("中指")
	} else if finger == 4 {
		fmt.Println("食指")
	} else {
		fmt.Println("大姆指")
	}

	// 可读性更强
	switch finger {
	case 1:
		fmt.Println("小姆指")
	case 2:
		fmt.Println("无名指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("食指")
	case 5:
		fmt.Println("大姆指")
	}

	// case 后支持多个参数
	switch n := 3; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	// case 后支持判断
	switch age := 20; {
	case age <= 18:
		fmt.Println("good good study")
	case age > 18 && age <= 60:
		fmt.Println("good good work")
	case age > 60:
		fmt.Println("good good play")
	default:
		fmt.Println(age)
	}

	// fallthrough
	switch {
	case false:
		fmt.Println("false1")
		fallthrough
	case true:
		fmt.Println("true1") // exe
		fallthrough          // 直接执行下面的case，不管是否为值
	case false:
		fmt.Println("false2") // exe
		fallthrough           // 直接执行下面的case，不管是否为值
	case true:
		fmt.Println("true2") // exe
		fallthrough          // 直接执行下面的case，不管是否为值
	case false:
		fmt.Println("false3") // exe
		fallthrough           // 直接执行下面的case，不管是否为值
	default:
		fmt.Println("default case") // exe
	}
}
