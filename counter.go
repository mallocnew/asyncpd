// Copyright 2024 GOTHAM Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)
package main

import (
	"fmt"
	"sync"
)

func main() {
	evenCh, oddCh := make(chan bool, 1), make(chan bool, 1)
	defer close(evenCh)
	defer close(oddCh)
	wg := sync.WaitGroup{}
	wg.Add(2)
	evenCh <- true
	go odd(oddCh, evenCh, &wg)
	go even(evenCh, oddCh, &wg)
	wg.Wait()
}

func even(evenCh, oddCh chan bool, wg *sync.WaitGroup) {
	for i := 2; i <= 10; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- true
	}
	wg.Done()
}

func odd(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i += 2 {
		<-evenCh
		fmt.Println(i)
		oddCh <- true
	}
	wg.Done()
}
