package main

import (
	"time"
	"fmt"
	"github.com/golangframework/map256"

	"strconv"
)
var c=make(chan string,2)
var count=30000000;
func main() {

	var m=make(map[string]int,count)
	var x=map256.Map256{}

	for i :=0; i <count ; i++ {
		x.Put([]byte(strconv.Itoa(i)),i)
		m[(strconv.Itoa(i))]=i
	}
	go initmap256(x)
	go initmap(m)



	fmt.Println(<-c)
	fmt.Println(<-c)
}

func initmap256(n map256.Map256)  {
	mtime := time.Now().UnixNano()
	for i :=0; i <count ; i++ {
		var _ interface{}=n.Get([]byte(strconv.Itoa(i)))

	}
	mtime1 := time.Now().UnixNano()
	var s=float64(mtime1-mtime)/float64(1*time.Second)
	c<-"map256:"+fmt.Sprint(s)
}

func initmap(n map[string]int)  {
	mtime := time.Now().UnixNano()
	for i :=0; i <count ; i++ {
		var _=n[(strconv.Itoa(i))]
	}
	mtime1 := time.Now().UnixNano()
	var s=float64(mtime1-mtime)/float64(1*time.Second)
	c<-"map:"+fmt.Sprint(s)
}