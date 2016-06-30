package main

import (
	"time"
	"fmt"
	"github.com/golangframework/map256"

	"strconv"
)


func main() {
	test(100)
	test(1000)
	test(10000)
	test(100000)
	test(1000000)
	test(10000000)
	test(20000000)
	test(30000000)
	test(40000000)




}
func test(count int)  {
	var c=make(chan string,2)
	var m=make(map[string]int,count)
	var x=map256.Map256{}

	for i :=0; i <count ; i++ {
		x.Put([]byte(strconv.Itoa(i)),i)
		m[(strconv.Itoa(i))]=i
	}
	go initmap256(x,count,c)
	go initmap(m,count,c)


	fmt.Print(strconv.Itoa(count)+"\t")
	fmt.Print(<-c)
	fmt.Println(<-c)
}
func initmap256(n map256.Map256,count int,ch chan string)  {
	mtime := time.Now().UnixNano()
	for i :=0; i <count ; i++ {
		var _ interface{}=n.Get([]byte(strconv.Itoa(i)))

	}
	mtime1 := time.Now().UnixNano()
	var s=float64(mtime1-mtime)/float64(1*time.Second)
	ch<-"map256\t"+fmt.Sprint(s)
}

func initmap(n map[string]int,count int,ch chan string)  {
	mtime := time.Now().UnixNano()
	for i :=0; i <count ; i++ {
		var _=n[(strconv.Itoa(i))]
	}
	mtime1 := time.Now().UnixNano()
	var s=float64(mtime1-mtime)/float64(1*time.Second)
	ch<-"\tmap\t"+fmt.Sprint(s)
}