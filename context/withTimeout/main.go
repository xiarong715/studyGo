package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) // 2s后超时
	defer cancel()

	go func() {
		time.Sleep(time.Second * 1) // 执行1s
		cancel()                    // 发送取消信号（关闭channel），可能在发送取消信号前就超时了
	}()

	select {
	case <-time.After(time.Second * 5): // 5s后才可能被执行
		fmt.Println("overslept")
	case <-ctx.Done(): // 接收信号
		fmt.Println(ctx.Err()) // "context canceled": 在预定时间内完成。  "context deadline exceeded": 在预定时间内没完成。
	}
}
