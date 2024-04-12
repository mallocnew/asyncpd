// Copyright 2024 GOTHAM Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	ch := make(chan int, 10)
	go sendData(ch)
	receiveData(ch)
}

func sendData(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed!")
			return
		}
		fmt.Println("Receive: ", v)
	}
}
