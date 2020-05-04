package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	recv(even, odd, quit)
}

func send(even, odd, quit chan<- int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}

func recv(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even:\t", v)
		case v := <-odd:
			fmt.Println("Odd:\t", v)
		case <-quit:
			return
		}
	}
}
