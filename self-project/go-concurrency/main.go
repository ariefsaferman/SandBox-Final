package main

import (
	"fmt"
	"time"
)

var channel = make(chan string)

func main() {
	// defer close(channel)
	// go sendData()

	// data := <-channel
	// fmt.Println(data)

	// time.Sleep(5 * time.Second)

	// go say("world")
	// say("hello")
	// fmt.Println("Hello")

	//buffer
	msg := make(chan int, 3)

	go func() {
		for {
			i := <-msg
			fmt.Println("Received data <-", i)
		}

	}()

	for i := 1; i < 9; i++ {
		msg <- i
		fmt.Println("Sent data ->", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func sendData() {
	time.Sleep(2 * time.Second)
	channel <- "Digitalent"
	fmt.Println("Done sending data to channelW")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
