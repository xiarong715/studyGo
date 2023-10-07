package main

import (
	"fmt"
	"time"
)

func go_worker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我的名字是", name)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "执行完毕")
}

func main() {
	go_worker("123")
	go_worker("456")
	for i := 0; i < 5; i++ {
		fmt.Println("我是main")
		time.Sleep(1 * time.Second)
	}
}
