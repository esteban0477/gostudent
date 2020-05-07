package main

import "fmt"

func main() {
	i := c()
	fmt.Println(i)
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}
