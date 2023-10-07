package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 发送消息
func Send(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- "hello"
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}

func Recv(ch <-chan string) {
	for d := range ch {
		fmt.Println(d)
	}
	wg.Done()
}

func main() {
	ch := make(chan string, 10)
	wg.Add(1)
	go Recv(ch)

	Send(ch)
	wg.Wait()
}
