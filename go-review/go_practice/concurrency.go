package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// ---------------------------- GoRoutine ----------------------------
	// f("direct")

	// go f("Goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	// ---------------------------- Channels ----------------------------
	// channels()

	//---------------------------- Channel Buffering ----------------------------
	// channelsBuffering()

	//---------------------------- Channel Synchronization ----------------------------
	// channelSynchronization()

}

func channels() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func channelsBuffering() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Print("Working....")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func channelSynchronization() {
	done := make(chan bool)
	go worker(done)

	<-done
}
