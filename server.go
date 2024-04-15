// Copyright 2024 GOTHAM Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print(r)
	fmt.Fprintf(w, "Hello\n")
	time.Sleep(10 * time.Second)
	log.Print("response")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	server := http.Server{Addr: ":8080", Handler: mux}
	log.Print("Server on: 8080")
	go server.ListenAndServe()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("ShutDown: ", err)
	} else {
		log.Print("ShutDown")
	}
}
