package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Println(a.Name, "eat food.")
}

func (a *Animal) Sleep() {
	fmt.Println(a.Name, "sleep.")
}

func (a *Animal) Shout() {
	fmt.Println(a.Name, "shout when it hungry.")
}

// 组合方式实现继承
type Cat struct {
	Animal
	Skill string
}

// 继承父类的属性，覆写父类的方法
func (c *Cat) Eat() {
	fmt.Println(c.Name, "eat food. 子类")
}

// 继承父类的属性，实现子类独有的方法
func (c *Cat) Ability() {
	fmt.Println(c.Name, c.Skill)
}

func NewCat(name string, skill string) *Cat {
	return &Cat{
		Animal: Animal{
			Name: name,
		},
		Skill: skill,
	}
}

func main() {
	cat := NewCat("zai zai", "catch mice")
	cat.Eat()     // 调用覆写的Eat方法
	cat.Sleep()   // 调用基类的Sleep方法
	cat.Shout()   // 调用基类的Shout方法
	cat.Ability() // 调用自己独有的Ability方法
}
