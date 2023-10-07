package main

import (
	"encoding/json"
	"fmt"
)

type custom struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      *int8  `json:"age,omitempty"`
	Status   string `json:"status"`
}

func json2struct(str string) error {
	c := &custom{}
	err := json.Unmarshal([]byte(str), c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if c.Age != nil {
		fmt.Println(*c.Age)
	}
	fmt.Println(c.Name, "--", c.Password, "--", c.Status)

	return nil
}

func main() {
	// 没写的字段解析出来是零值
	// 不需要详细解析的字段就写成json.RawMessage字段
	json := `{"name":"xiarong","password":"123456","status":"100","payload":{"u":"xiarong"}}`
	_ = json2struct(json)
}
