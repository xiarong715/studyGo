package main

import (
	"encoding/json"
	"fmt"
)

type replyStruct struct {
	Result string
	Status int8
	Token  string
}

func replyMessage(ok bool, replystruct replyStruct) ([]byte, error) {
	reply := ""
	if ok {
		reply = "ok"
	} else {
		reply = "failed"
	}

	stat := new(int8)
	if replystruct.Status == -1 {
		stat = nil
	} else {
		*stat = replystruct.Status
	}

	return json.Marshal(struct {
		Reply  string `json:"reply"`
		Token  string `json:"token,omitempty"` // omitempty的作用：当为零值时，不会出现在json中
		Result string `json:"result,omitempty"`
		Status *int8  `json:"status,omitempty"`
	}{
		Reply:  reply,
		Token:  replystruct.Token,
		Status: stat,
		Result: replystruct.Result,
	})
}

func replyMessage2(ok bool, msg []string) []byte {
	r := "failed"
	if ok {
		r = "ok"
	}

	var reply []byte
	var err error

	if len(msg) > 1 {
		reply, err = json.Marshal(struct {
			Reply  string   `json:"reply"`
			Result []string `json:"result"`
		}{
			Reply:  r,
			Result: msg,
		})
	} else {
		if len(msg) == 0 {
			return []byte(`{"reply":"failed","result":"` + "" + `"}`)
		}
		reply, err = json.Marshal(struct {
			Reply  string `json:"reply"`
			Result string `json:"result"`
		}{
			Reply:  r,
			Result: msg[0],
		})
	}
	if err != nil {
		return []byte(`{"reply":"failed","result":"` + err.Error() + `"}`)
	}
	return reply
}

func main() {
	r := replyStruct{Result: "username and password not match", Status: 1, Token: "12345"}
	reply, err := replyMessage(false, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(reply))

	fmt.Println(string(replyMessage2(false, []string{"hello", "world"})))
}
