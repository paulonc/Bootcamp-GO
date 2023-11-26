package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
	}
}
func pong(c chan string) {
	for {
		c <- "pong"
	}
}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	c := make(chan string)

	go ping(c)
	go imprimir(c)
	go pong(c)

	var entrada string
	fmt.Scanln(&entrada)
}
