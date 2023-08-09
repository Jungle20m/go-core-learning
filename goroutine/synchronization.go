package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayHello(times int) {
	for i := 0; i < times; i++ {
		log.Printf("hello, %d\n", i)
		time.Sleep(time.Second)
	}

	wg.Done()
}

func SayHi(times int) {
	for i := 0; i < times; i++ {
		log.Printf("hi, %d\n", i)
		time.Sleep(time.Second)
	}

	wg.Done()
}

func Synchronization1() {
	wg.Add(2)

	go SayHi(10)
	go SayHello(10)

	wg.Wait()
	log.Println("done")
}
