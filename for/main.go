package main

import "fmt"

// 1 * 1 = 1
// 1 * 2 = 2   2 * 2 = 4
// 1 * 3 = 3   2 * 3 = 6   3 * 3 = 9

func main() {
	for i := 1; i <= 9; i++ { // 注意 i 的作用域，在第一层循环内
		for j := 1; j <= i; j++ { // 注意 j 的作用域，在第二层循环内
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
