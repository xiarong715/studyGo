package main

import "fmt"

type USB interface {
	start()
	stop()
}

type Camera struct {
}

func (c *Camera) start() {
	fmt.Println("camera start...")
}

func (c Camera) stop() {
	fmt.Println("camera stop...")
}

type Phone struct {
}

func (p Phone) start() {
	fmt.Println("phone start...")
}
func (p Phone) stop() {
	fmt.Println("phone stop...")
}

type Computer struct {
}

// 接口变量没有指针类型
// 可以将实现该接口的指针变量赋值给接口变量
// 接收者为指针类型实现的接口，只能用指针变量调用接口方法。
// 接收者为值类型实现的接口，可用指针变量调用也可用值变量调用接口方法。
// 接口变量被赋指针类型变量，则接口变量能同时调用接收者为指针类型和值类型的接口方法。
// 接口变量被赋值类型变量，则接口变量只能调用接收者为指针类型的接口方法。
// 接口变量被赋实现接口的结构体值变量或指针变量，用接口变量调用接口方法，实现多态。
func (computer Computer) Working(usb USB) {
	usb.start()
	usb.stop()
}

// 接收者：值类型 指针类型
// 调用者：值类型 指针类型
// 不通过接口变量调用方法时，调用者为值类型或指针类型，可调用接收者为值类型或接收者为指针类型的方法。
// 通过接口变量调用方法时，赋给接口变量的是指针类型变量，则接口变量能调用接收者为指针类型的方法，也能调用接收者为值类型的方法。
// 						 赋给接口变量的是值类型变量，  则接口变量只可调用接收者为值类型的方法。

func main() {
	var usb USB
	c := &Camera{}
	p := Phone{}
	usb = c     // 赋给接口变量的是指针类型变量
	usb.start() // 则接口变量可调用接收者为指针类型的方法
	usb.stop()  // 也可调用接收者为值类型的方法。
	usb = p     // 赋给接口变量的是值类型变量
	usb.start() // 则接口变量只可调用接收者为值类型的方法。
	usb.stop()  // 则接口变量只可调用接收者为值类型的方法。

	computer := Computer{}
	computer.Working(c)
	computer.Working(p)
}
