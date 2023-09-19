package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

var muser atomic.Value
var mtoken atomic.Value
var mu sync.Mutex
var gUserFile *os.File
var gTokenFile *os.File

// var gChData chan chdata

type UserManagerJSON struct {
	Users  map[string]User   `json:"users"`
	Tokens map[string]string `json:"tokens"`
	Secret string            `json:"secret"`
}

type User struct {
	Password string            `json:"password"`
	Email    string            `json:"email"`
	Salt     string            `json:"salt"`
	Info     map[string]string `json:"info"`
}

func init() {
	muser.Store(make(map[string]User))
	mtoken.Store(make(map[string]string))

	var err error
	gUserFile, err = os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	gTokenFile, err = os.OpenFile("tokens.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// gChData = make(chan chdata, 10)
}

func usersRead(name string) User {
	m1 := muser.Load().(map[string]User)
	return m1[name]
}

func usersInsert(name string, newuser User) error {
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

func usersGet() map[string]User {
	m1 := muser.Load().(map[string]User)
	m2 := make(map[string]User)
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

func tokensRead(appname string) string {
	m1 := mtoken.Load().(map[string]string)
	return m1[appname]
}

func tokensInsert(appname string, token string) error {
	mu.Lock()
	defer mu.Unlock()
	m1 := mtoken.Load().(map[string]string)
	m2 := make(map[string]string)
	for k, v := range m1 {
		m2[k] = v
	}
	m2[appname] = token
	mtoken.Store(m2)
	return nil
}

func tokensGet() map[string]string {
	m1 := mtoken.Load().(map[string]string)
	m2 := make(map[string]string)
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

func SaveJSON11(tag string, ch chan<- chdata) error {
	switch tag {
	case "USERS":
		users := usersGet()
		data, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return SaveJSON(tag, data, ch)
	case "TOKENS":
		tokens := tokensGet()
		data, err := json.Marshal(tokens)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return SaveJSON(tag, data, ch)
	default:
		fmt.Println("tag is err")
		return errors.New("tag is err")
	}
}

func SaveJSON(tag string, data []byte, ch chan<- chdata) error {
	ch <- chdata{Tag: tag, Data: data}
	return nil
}

func SaveToFile(data chdata) error {
	switch data.Tag {
	case "USERS":
		{
			return Save(gUserFile, data.Data)
		}
	case "TOKENS":
		{
			return Save(gTokenFile, data.Data)
		}
	default:
		return errors.New("tag is err")
	}
}

func Save(file *os.File, data []byte) error {
	file.Truncate(0)
	_, err := file.Write(data)
	if err != nil {
		return err
	}
	return err
}

// tag string, data []byte
type chdata struct {
	Tag  string
	Data []byte
}

func RoutingSaveToFile(ch <-chan chdata) { // 只能接收
	for data := range ch { // for { select {}}
		fmt.Println("routing save to file")
		SaveToFile(data)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	usersInsert("lsy", User{Password: "12345", Salt: "abcdef", Email: "lsy@ipanel.cn"})
	user := usersRead("lsy")
	fmt.Println(user.Email, user.Password, user.Salt)

	tokensInsert("gogs-user", "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy11c2VyIn0.BAKHMvBaQ5g7y1khYjWAX3lARKeNQ5TAx4yuHziDKf9O_gQGnkaDR9DOFoKG8P1KyVdLtbR-A_Ynl-boVzlUNg")
	tokensInsert("gogs-code", "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZ29ncy1jb2RlIn0.P5lHSe6cCGxPJTRDDEYYtR27FyaStm6ua3_210Tq_Hn5wnBTGt8yBvWlaSgPiPzwrPXO9U9vuXJ3LS0MT4RVOA")
	token := tokensRead("gogs-user")
	fmt.Println(token)

	users := usersGet()
	tokens := tokensGet()
	fmt.Println(users)
	fmt.Println(tokens)

	gChData := make(chan chdata, 10)

	SaveJSON11("USERS", gChData)
	SaveJSON11("TOKENS", gChData)

	wg.Add(1)
	go RoutingSaveToFile(gChData)

	wg.Wait()
}
