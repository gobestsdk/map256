package main

import (
	"github.com/golangframework/map256"
	"fmt"
)

func main() {
	m:=map256.Map256{}
	m.Put([]byte("1"),"我是李鹏")

	m.Put([]byte("12"),"我是李鹏爸爸")
	m.Put([]byte("4"),true)
	m.Put([]byte("5"),324)

	fmt.Print(m.String())
	m.Del([]byte("1"))


}
