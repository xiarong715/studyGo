package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func cMap() map[string]string {
	m := make(map[string]string)
	for i := 0; i < 100000; i++ {
		m[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	return m
}

func readMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
	wg.Done()
}

func writeMap(m map[string]string) {
	for k, v := range m {
		i, _ := strconv.Atoi(v)
		m[k] = strconv.Itoa(i + 1)
	}
	wg.Done()
}

// -------------------------------------------------------------------
var muser atomic.Value
var mu sync.Mutex

type User struct {
	Name  string
	Email string
}

func init() {
	muser.Store(make(map[string]User))
}

func getUser(name string) User {
	m := muser.Load().(map[string]User)
	return m[name]
}

func setUser(name string, newuser User) error {
	mu.Lock()
	defer mu.Unlock()
	m1 := muser.Load().(map[string]User)
	m2 := make(map[string]User)
	for k, v := range m1 {
		m2[k] = v
	}
	m2[name] = newuser
	muser.Store(m2)
	return nil
}

func getUsers() map[string]User {
	m1 := muser.Load().(map[string]User)
	m2 := make(map[string]User)
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

func main() {
	// m := cMap()
	// wg.Add(2)
	// go writeMap(m)
	// go readMap(m)

	wg.Add(2)
	go func() {
		for i := 0; i < 5000; i++ {
			setUser("user"+strconv.Itoa(i), User{Name: strconv.Itoa(i), Email: strconv.Itoa(i) + "@ipanel.cn"})
		}
		wg.Done()
	}()

	// go func() {
	// 	for i := 0; i < 5000; i++ {
	// 		user := getUser("user" + strconv.Itoa(i))
	// 		fmt.Println(user.Name, "---", user.Email)
	// 	}
	// 	wg.Done()
	// }()

	go func() {
		for i := 0; i < 5000; i++ {
			getUsers()
		}
		fmt.Println(getUsers())
		wg.Done()
	}()

	wg.Wait()

	// 不能多线程写
	// 不能多线程读写
	// 可以多线程读
}
