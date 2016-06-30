package main

import (
	"time"
	"fmt"
	"github.com/golangframework/map256"

	"strconv"
)

func main() {
	testinsert(100)
	testinsert(1000)
	testinsert(10000)
	testinsert(100000)
	testinsert(1000000)
	testinsert(10000000)
	testinsert(20000000)
	testinsert(30000000)
	testinsert(40000000)
}
func testinsert(count int){
	var c=make(chan string,2)
	go insertmap(count,c)
	go insertmap256(count,c)
	fmt.Print(strconv.Itoa(count)+"\t")
	fmt.Print(<-c)
	fmt.Println(<-c)
}
func insertmap(n int,ch chan string)  {
	mtime := time.Now().UnixNano()
	var m=make(map[string]int,n )
	for i :=0; i <n ; i++ {
		m[(strconv.Itoa(i))]=i
	}
	mtime1 := time.Now().UnixNano()

	var s=float64(mtime1-mtime)/float64(1*time.Second)

	ch<-"\tmap\t"+fmt.Sprint(s)
}

func insertmap256(n int,ch chan string)  {

	ttime := time.Now().UnixNano()
	var x=map256.Map256{}
	for i :=0; i <n ; i++ {
		x.Put([]byte(strconv.Itoa(i)),i)
	}
	ttime1 := time.Now().UnixNano()
	var s=float64(ttime1-ttime)/float64(1*time.Second)

	ch<-"\tmap256\t"+fmt.Sprint(s)
}