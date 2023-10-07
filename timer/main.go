package main

import (
	"fmt"
	"time"
)

// 看懂，默写

type H struct {
	t time.Ticker
}

func (h *H) ticker() {
	for range h.t.C {
		fmt.Println("ticker")
	}
}

func main() {
	h := &H{t: *time.NewTicker(time.Second * 1)}
	go h.ticker()
	time.Sleep(time.Second * 20)
}
