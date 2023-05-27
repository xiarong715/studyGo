package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go working(ctx)

	time.Sleep(time.Second * 1)
	cancelFunc()

	time.Sleep(time.Second * 1) // 确保协程接受到取消信号
}

func working(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("workdone")
			return
		default:
			fmt.Println("working")
		}
	}
}
