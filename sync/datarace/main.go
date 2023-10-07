package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		// 一个线程对x的修改，会影响另一个线程中x的值。如线程A，拿到的x为1000，执行权给了线程B，线程B继续对x操作，加到1150。
		// 然后执行权给了线程C，对x继续加到1220。
		// 执行权限返回到线程A，线程A从1000开始对x加1操作，线程A的操作使用x的值错乱，线程都运行完后，x的值不正确。
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	fmt.Println("start.")
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
	fmt.Println("end.")
}
