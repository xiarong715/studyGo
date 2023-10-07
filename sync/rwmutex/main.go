package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex
var rwmu sync.RWMutex
var x int64

func readWithLock() {
	mu.Lock()
	time.Sleep(1 * time.Millisecond)
	mu.Unlock()
	wg.Done()
}

func writeWithLock() {
	mu.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	mu.Unlock()
	wg.Done()
}

func readWithRWLock() {
	rwmu.RLock()
	time.Sleep(1 * time.Millisecond)
	rwmu.RUnlock()
	wg.Done()
}

func writeWithRWLock() {
	rwmu.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwmu.Unlock()
	wg.Done()
}

func do(rf, wf func(), rc, wc int) {
	start := time.Now()
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}
	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x: %v cost: %v\n", x, cost)
}

func main() {
	do(readWithLock, writeWithLock, 1000, 10)     // 读多，写少
	do(readWithRWLock, writeWithRWLock, 1000, 10) // 读多，写少
}
