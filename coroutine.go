// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func SleepLoop(f *atomic.Bool, w *sync.WaitGroup) {
	fmt.Printf("function[pid: %d, tid: %d] Begin\n", os.Getpid(), GoID())
	for f.Load() == false {
		fmt.Printf("function[pid: %d, tid: %d]: time %s\n", os.Getpid(), GoID(), time.Now().String())
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("function[pid: %d, tid: %d] Over\n", os.Getpid(), GoID())
	w.Done()
}

func main() {
	fmt.Printf("Main function, pid: %d, tid: %d\n", os.Getpid(), GoID())
	f := atomic.Bool{}
	var w sync.WaitGroup
	w.Add(3)
	go SleepLoop(&f, &w)
	go SleepLoop(&f, &w)
	go SleepLoop(&f, &w)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	f.Store(true)
	fmt.Println("Got signal: ", s)
	w.Wait()
}
