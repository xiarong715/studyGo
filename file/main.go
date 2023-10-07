package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("hello.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("hello file"))
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := file.Seek(0, io.SeekStart) // 相对于文件首部移动0个位置
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Truncate("hello.txt", 0) // clear
	fmt.Println("n = ", n)

	file.WriteAt([]byte("I love zaizai."), n) // 在文件首部写内容
}
