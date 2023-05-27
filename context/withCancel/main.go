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
	cancelFunc() // 发送取消信号（关闭channel）

	time.Sleep(time.Second * 1) // 确保协程接收到取消信号
}

func working(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // 接收取消信号
			fmt.Println("workdone")
			return
		default:
			fmt.Println("working")
		}
	}
}
