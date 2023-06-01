package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func sortMap() {
	m2 := make(map[string]int, 100)
	rand.New(rand.NewSource(time.Now().UnixNano())) // 把时间戳设置为时间种子
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		m2[key] = value
	}
	keys := make([]string, 0)
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Strings(keys) // 对key排序
	for _, key := range keys {
		fmt.Println(key, m2[key])
	}
}

func main() {
	// 声明
	var m map[string]string // nil
	// 初始化
	m = make(map[string]string, 10) // 要估算好容量，避免程序在运行期间动态扩容，影响性能
	m["hello"] = "world"
	m["map"] = "map"
	fmt.Println(m)

	// 取值
	fmt.Println(m["hello"])
	fmt.Println(m["map"])
	fmt.Println(m["lsy"]) // ""

	// 取值判断
	value, ok := m["lsy"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

	// 遍历取值
	// 取 键、值
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 只取键
	for k := range m {
		fmt.Println(k)
	}
	// 只取值
	for _, v := range m {
		fmt.Println(v)
	}

	// 删除
	delete(m, "lsy") // 删除不存在的键也不报错
	delete(m, "map")

	// 排序
	sortMap()
	// 应用实践
	// 如何生切片类型的map
	ms := make(map[string][]int, 5)
	ms["hello"] = []int{1, 2, 3}
	ms["lsy"] = []int{2, 3, 4}
	ms["key"] = []int{4, 5, 6}
	fmt.Printf("%T\n", ms)
	fmt.Printf("%v\n", ms)

	// 思考题
	// 只能通过make函数初始化内存，不能赋值，其值为类型的零值。
}
