package main

import (
	"github.com/golangframework/map256"
	"strconv"
)

func main() {
	var m=map256.Map256{}
	var ch =make(chan int,2)
	go func() {
		for i:=0;i<100;i+=2{
			m.Put([]byte(strconv.Itoa(i)),"go1:"+strconv.Itoa(i))
		}
		ch<-1
	}()
	go func() {
		for i:=1;i<100;i+=2{
			m.Put([]byte(strconv.Itoa(i)),"go2:"+strconv.Itoa(i+1))
		}
		ch<-1
	}()
	<-ch
	<-ch
	print(m.String())
}
