package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		// 加互斥锁，避免数据错乱。互斥锁保证同一时间只有一个线程进入临界区，这样确保数据的正确性。
		// 当一个线程进入临界区时，会加锁，其他的线程会等待锁的释放，多个线程等待同一锁时，会随机唤醒等待的线程。
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
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
