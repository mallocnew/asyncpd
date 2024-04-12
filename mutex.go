// Copyright 2024 GOTHAM Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)
package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	x = x + 1
	lock.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Hello World!")
	var w sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &lock)
	}
	w.Wait()
	fmt.Println("val of x: ", x)
}
