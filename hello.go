package main

import (
	"fmt"
	"runtime"
)

var c chan bool

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c = make(chan bool)
	str := "0%"
	fmt.Println("开始...", str)
	go modify(&str)
	<-c
	fmt.Println("完成!", str)
}

func modify(ret *string) {
	*ret = "50%"
	fmt.Println("升级中", *ret)
	c <- true
	*ret = "100%"
}
