package main

import (
	"encoding/json"
	"fmt"
)

// 切片是引用类型。
// 引用类型：引用类型的变量中存的是地址，地址指向堆。
// Go中没有引用传递，只有值传递，变量的值为真实的值或为一个地址。
// Go中的引用类型：指针、map、slice、接口、函数。
// 在切片引用的底层数组中，从切片的第1个元素开始到数组的最后一个元素的长度就是切片的长度。

// 值传递
func modifySlice(s []int) {
	s = append(s, 4)
	s[0] = 10 // 修改的是切片最新指向的数组
} // 所以：只要对切片追加元素，都要返回最新的切片，避免出现bug

func modifySlice2(s *[]int) {
	*s = append(*s, 4)
	(*s)[0] = 10
}

func main() {
	s1 := []string{"hello", "world", "lsy"}
	fmt.Printf("len(s1) = %v cap(s1) = %v\n", len(s1), cap(s1)) // len(s1) = 3 cap(s1) = 3
	fmt.Printf("%p\n", s1)                                      // 0xc0000a0150

	// 追加一个元素。
	// 切片容量只是3，且满了，因此会申请更大长度的数组，作为切片的底层数组，把旧数组中的数据拷贝到新数组中，并放入追加的元素。
	// 旧数组失去引用，等待GC回收。
	s1 = append(s1, "cap")                                      // 重新分配底层数组
	fmt.Printf("len(s1) = %v cap(s1) = %v\n", len(s1), cap(s1)) // len(s1) = 4 cap(s1) = 6
	fmt.Printf("%p\n", s1)                                      // 0xc0000b0120

	fmt.Printf("\n-------------------------\n\n")

	// 切片长度
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}                         // 定义一个数组
	s2 := arr[2:5]                                               // [3,4,5]
	fmt.Printf("len(s2) = %v, cap(s2) = %v\n", len(s2), cap(s2)) // len(s2) = 3 cap(s2) = 5
	s3 := arr[3:]                                                // [4,5,6,7]
	fmt.Printf("len(s3) = %v, cap(s3) = %v\n", len(s3), cap(s3)) // len(s3) = 4 cap(s3) = 4
	s4 := arr[:]                                                 // [1,2,3,4,5,6,7]
	fmt.Printf("len(s4) = %v, cap(s4) = %v\n", len(s4), cap(s4)) // len(s4) = 7 cap(s4) = 7
	s5 := s4[5:]                                                 // 对切片再切片 [6,7]
	fmt.Printf("len(s5) = %v, cap(s5) = %v\n", len(s5), cap(s5)) // len(s5) = 2 cap(s5) = 2

	fmt.Printf("\n-------------------------\n\n")

	// 用操作説明切片是引用类型
	arr[4] = 555            // 修改底层数组的值
	fmt.Println(s2, s3, s4) // 引用该数组的切片都受到影响 [3,4,555] [4,555,6,7] [1,2,3,4,555,6,7]

	fmt.Printf("\n-------------------------\n\n")
	// make 生成切片, append 切片的追加, copy 切片的复制
	// 分配底层数组的规则：每次分配置的容量都是前一次的2倍
	s6 := make([]int, 0)
	fmt.Printf("len(s6) = %v cap(s6) = %v\n", len(s6), cap(s6)) // len(s6) = 0 cap(s6) = 0
	s6 = append(s6, 2, 3, 4, 5)                                 // 首次追加时量体裁衣，之后分配时是之前的2倍。
	fmt.Printf("len(s6) = %v cap(s6) = %v\n", len(s6), cap(s6)) // len(s6) = 2 cap(s6) = 4
	s6 = append(s6, 7)
	fmt.Printf("len(s6) = %v cap(s6) = %v\n", len(s6), cap(s6)) // len(s6) = 3 cap(s6) = 8

	s7 := make([]int, 2, 5)
	fmt.Printf("len(s7) = %v cap(s7) = %v\n", len(s7), cap(s7)) // len(s7) = 2 cap(s7) = 5
	fmt.Println(s7)
	s7 = append(s7, 3, 4, 5, 6)                                 // 超过了容量5，重新分配底层数组
	fmt.Printf("len(s7) = %v cap(s7) = %v\n", len(s7), cap(s7)) // len(s7) = 2 cap(s7) = 5*2
	fmt.Println(s7)

	// slice
	s8 := []int{1, 2, 3}
	s9 := s8
	s8[0] = 10
	fmt.Println(s8, s9) // s8 s9指同一个底层数组 [10,2,3] [10,2,3]

	// array
	s10 := [3]int{1, 2, 3}
	s11 := s10 // 数组是值类型，赋值时会给s11分配空间，并拷贝值给s11
	s10[0] = 1000
	fmt.Println(s10, s11) // 两数组没关联，修改一个数组的值，另一个数组不会发生改变 [1000,2,3] [1,2,3]

	// slice的遍历
	// for i 或 for range
	s12 := []int{7, 8, 9}
	for i := 0; i < len(s12); i++ {
		fmt.Print(s12[i], " ")
	}
	fmt.Println()

	for _, v := range s12 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// append追加
	s13 := []int{1, 2, 3}
	fmt.Printf("len(s13) = %v cap(s13) = %v\n", len(s13), cap(s13)) // len(s13) = 3 cap(s13) = 3
	s13 = append(s13, 4)                                            // 追加一个元素
	fmt.Printf("len(s13) = %v cap(s13) = %v\n", len(s13), cap(s13)) // len(s13) = 4 cap(s13) = 6    // 当容量满，追加元素时，分配的容量是当前的2倍

	s14 := []int{5, 6, 7}
	s14 = append(s14, s13...) //追加一个切片，注意有3个点
	// s14 = append(s14, []int{1,2,3}...)  // 这种形式也可以
	fmt.Println(s14)

	// copy 复制切片
	// 因为切片是引用类型，因为通过复制，切断两切片的关联
	s15 := []int{1, 2, 3}
	s16 := s15 // 引用同一个内存地址，s15和s16有关联关系
	fmt.Println(s16)

	s17 := []int{2, 3, 4}
	s18 := make([]int, 3)
	copy(s18, s17) // 把s17的内容复制到s18中，s17和s18没关联关系，这也是copy的意义

	// 删除元素
	s19 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("del before: ", s19)
	s19 = append(s19[:3], s19[4:]...) // 删除下标为3的元素
	fmt.Println("del after: ", s19)

	// 数组转切片
	s20 := [...]int{1, 2, 3, 4, 5} // 数组
	s21 := s20[:]                  // 对数组切片
	fmt.Println(s21)
	fmt.Printf("s20 type: %T, s21 type: %T\n", s20, s21) // s20 type: [5]int, s21 type: []int

	// exam
	s22 := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		s22 = append(s22, i)
	}
	fmt.Println(s22)                                   // 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9
	fmt.Printf("%p %p %p\n", s22[:], s22[1:], s22[2:]) // 0xc0000c2000 0xc0000c2008 0xc0000c2010  每个元素占8个字节，元素之间的地址相邻

	// 直接定义与make定义的区别
	var s []int          // 没给切片分配底层数组，指针字段指向nil
	ss := make([]int, 0) // 分配了底层数组，只不过是包含0个元素的内存空间
	fmt.Println(s, ss)

	res1, _ := json.Marshal(s)
	res2, _ := json.Marshal(ss)
	res3, _ := json.Marshal([]int{1, 2, 3})
	fmt.Println(string(res1)) // 注意：是null
	fmt.Println(string(res2)) // []
	fmt.Println(string(res3)) // [1, 2, 3]

	sss := []int{1, 2, 3}
	modifySlice(sss) // 在函数中追加元素后，扩容了，底层数组改变了。  但sss指向的底层数组没变。因此只要对切片做append的操作，都要返回新的切片
	fmt.Println(sss) // [1 2 3]

	modifySlice2(&sss) // 传切片的地址，可以保证正确修改切片的值。和C语言类似，传指针，对正确的改变切片的值。
	fmt.Println(sss)   // [10 2 3 4]
}
