package main

import "fmt"

type Action interface {
	Eat()
	Sleep()
}

type Person struct {
	Name string
	Age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int8 {
	return p.Age
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) SetAge(age int8) {
	p.Age = age
}

func (p *Person) Eat() {
	fmt.Println(p.Name, "eat.")
}

func (p *Person) Sleep() {
	fmt.Println(p.Name, "sleep.")
}

// 内嵌结构体，实现继承
// 能共享基类的字段
type Student struct {
	Person
	Number string
}

func NewStudent(name string, age int8, num string) *Student {
	return &Student{
		Person: *NewPerson(name, age),
		Number: num,
	}
}

func (s Student) GetNum() string {
	return s.Number
}

func (s *Student) SetNum(num string) {
	s.Number = num
}

// 内嵌接口，实现继承
// 只能调用基类的接口方法，但可实现多态
// 不能共享基类字段
type Engineer struct {
	Action
	Name  string
	Age   int8
	Kinds string
}

func NewEngineer(action Action, name string, age int8, kinds string) *Engineer {
	return &Engineer{
		Action: action,
		Name:   name,
		Age:    age,
		Kinds:  kinds,
	}
}

// 覆写基类方法
func (e Engineer) Eat() {
	fmt.Println(e.Name, "eat.")
}

func (e Engineer) Sleep() {
	fmt.Println(e.Name, "sleep.")
}

func main() {
	per := NewPerson("zai zai", 2)
	fmt.Println(per.GetName(), per.GetAge())
	fmt.Println("--------------------------")

	per.SetName("haohao")
	per.SetAge(30)
	fmt.Println(per.GetName(), per.GetAge())
	per.Eat()
	per.Sleep()

	var action Action = per // 结构体对象赋给接口变量
	action.Eat()            // 接口变量调用方法
	action.Sleep()
	fmt.Println("--------------------------")

	stu := NewStudent("lingling", 30, "001")
	fmt.Println(stu.GetName(), stu.GetAge(), stu.GetNum())
	stu.Eat()
	stu.Sleep()
	fmt.Println("--------------------------")

	// 实现多态：接口和接口类型切片
	actions := []Action{per, stu}
	for _, action := range actions {
		action.Eat()
		action.Sleep()
	}

	fmt.Println("--------------------------")

	engineer := NewEngineer(per, "xr", 30, "IT")
	engineer.Eat()
	engineer.Sleep()
}
