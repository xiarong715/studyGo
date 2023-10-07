package main

import "fmt"

func worker(c chan int) {
	num := <-c //读取管道中的数据，并输出
	fmt.Println("接收到参数c:", num)
}

func main() { //channel的创建,需要执行管道数据的类型，我们这里是
	c := make(chan int) //开辟一个协程 去执行worker函数
	go worker(c)
	c <- 2 //往管道中写入2
	fmt.Println("main")
}
