package main

import (
	"github.com/golangframework/map256"
	"fmt"
)

func main() {
	m:=map256.Map256{}

	m.Put([]byte("123"),"我是李鹏爸爸")
	m.Put([]byte("1234"),true)
	m.Put([]byte("123456"),324)

	fmt.Println(m.String())
	fmt.Println(m.Len())
	m.Del([]byte("1"))


}
