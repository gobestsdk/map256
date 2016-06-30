package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var c=make(chan string,2)
	go doingtmap(100000000,c)

	fmt.Print(strconv.Itoa(100000000)+"\t")
	fmt.Print(<-c)

}
func doingtmap(n int,ch chan string)  {
	var m=make(map[string]int,n )
	mtime := time.Now().UnixNano()
	for i :=0; i <n ; i++ {
		m[(strconv.Itoa(i))]=i
	}
	mtime1 := time.Now().UnixNano()

	var s=float64(mtime1-mtime)/float64(1*time.Second)

	ch<-"\tmap\t"+fmt.Sprint(s)
}

