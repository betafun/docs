package main

import (
	"encoding/json"
	"fmt"
)

type MyJsonName struct {
	ErrMsg  string `json:"err_msg"`
	RetCode int    `json:"ret_code"`
}

func main() {
       // exapme 1
       var s1 MyJsonName
       s1.ErrMsg = "hello s1"
       s1.RetCode  = 200
       a, _ := json.Marshal(s1)
       fmt.Println(string(a))

       s2 := &MyJsonName {
	"hello s2",
	300, 
        }
	b, _ := json.Marshal(s2)
       fmt.Println(string(b))

       s3 := MyJsonName{ RetCode:400,ErrMsg:`hello s3`}
       c, _ := json.Marshal(s3)
       fmt.Println(string(c))
}
