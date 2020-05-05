package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOut(c1, c2)

	for Printer := range c2 {
		fmt.Println(Printer)
	}

}

// Function populate as it's name show, it populate the channel with
// the Fibonnacci numbers
func populate(c chan int) {
	const rangeFibo = 20
	actual := 0
	lastone := 0
	lasttwo := 1

	for i := 0; i < rangeFibo; i++ {
		actual = lastone + lasttwo
		c <- actual
		lasttwo = lastone
		lastone = actual
	}
	close(c)
}

func fanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for value := range c1 {
		wg.Add(1)
		go func(valueToWait int) {
			c2 <- waitSomeTime(valueToWait)
			wg.Done()
		}(value)
	}
	wg.Wait()
	close(c2)
}

func waitSomeTime(valueToWait int) int {
	time.Sleep(time.Second * time.Duration(valueToWait))
	return valueToWait
}
