package main

import "log"

type reply struct {
	Name string
}

func getUserList() interface{} {

	r := reply{Name: "lsy"}

	return &r
}

func main() {
	v := getUserList()
	ret, ok := v.(*reply)
	if !ok {
		log.Println("assertion err")
		return
	}
	log.Println(*ret)
}
