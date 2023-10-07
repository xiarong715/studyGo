package main

import "fmt"

// 接口。相当于抽象类
type Animal interface {
	Eat()
	Sleep()
	Shout()
}

// 实现Animal接口，使AnimalImpl具有Animal的特征。AnimalImpl相当于基类。
type AnimalImpl struct {
	Name string
}

// 构造函数
func NewAnimalImpl(name string) *AnimalImpl {
	return &AnimalImpl{
		Name: name,
	}
}

func (a *AnimalImpl) Eat() {
	fmt.Println(a.Name, " eat fish.")
}

func (a *AnimalImpl) Sleep() {
	fmt.Println(a.Name, " sleep in daytime.")
}

func (a *AnimalImpl) Shout() {
	fmt.Println(a.Name, " shout when it hungry.")
}

// 继承Animal接口
type Cat struct {
	Animal        // 内嵌接口，无法继承基类的属性字段，但能实现方法的多态，能覆写基类的方法。
	Name   string // 需要加上自己的字段。
}

// 覆写基类AnimalImpl的方法
func (c *Cat) Eat() {
	fmt.Println(c.Name, "eat fish.")
}

func (c *Cat) Sleep() {
	fmt.Println(c.Name, "sleep in daytime.")
}

// func (c *Cat) Shout() {
// 	fmt.Println(c.Name, "shout when it hungry.")
// }

func NewCat(a Animal, name string) *Cat {
	return &Cat{
		Animal: a,
		Name:   name,
	}
}

// 不修改结构体，或结构体占空间不大时，可把结构体方的接收者设置为值类型。

func main() {
	animal := NewAnimalImpl("miaomiao")
	// cat := &Cat{
	// 	Animal: animal, // 初始化基类
	// 	Name:   "zai zai",
	// }
	cat := NewCat(animal, "zai zai")
	cat.Eat() // 调用方法时，首先寻找本结构体中是否有该方法，没有的话，会在嵌套的结构体中寻找方法。通过在外层结构体中实现相同函数名的方法，达到覆写的目的。
	cat.Sleep()
	cat.Shout() // 派生类没有覆写基类的Shout方法。因此调用基类的Shout方法。
}
