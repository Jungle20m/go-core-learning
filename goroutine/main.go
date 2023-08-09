package main

import (
	"fmt"
	"runtime"
	"time"
)

func subProcess() {
	fmt.Println("sub process started")
	time.Sleep(time.Second * 3)
	fmt.Println("sub process stopped")
}

func mainProcess() {
	fmt.Println("main process started")

	go subProcess()

	time.Sleep(time.Second)
	fmt.Println("main process stopped")
}

func main() {
	runtime.GOMAXPROCS(6)
	//fmt.Println()
	fmt.Println(runtime.NumCPU())
	//Synchronization1()
}
