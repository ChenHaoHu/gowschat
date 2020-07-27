package main

import "fmt"

func main() {
	chan1 := make(chan int, 3)

	chan1 <- 1

	chan1 <- 2
	chan1 <- 2

	fmt.Println(<-chan1)

}
