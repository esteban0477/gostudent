package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	c := make(chan int)
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())

	// When 42 goes to the channel, we need another routine to receive the data at the same time,
	// Otherwise, this will block the execution of the code. In essence, When we're sengind data
	// to a channel, we need another routine to receive the data. For the next chunck of code, we
	// are creating another routine with a literal (anonymous) function, and we're sending the
	// interger value to the channel. In parallel, the Main routine will pull off 42 from the
	// channel and will print it out to the stdout

	go func() {
		c <- 42
	}()

	time.Sleep(time.Second)
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
	fmt.Println(<-c)
}
