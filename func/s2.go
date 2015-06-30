package main

import (
	"fmt"
)


func main() {
	c := make(chan int)
	
	go func(i int) {
		c <- i*i
	}(3)

	x := <-c
	fmt.Println(x)
	fmt.Println("Terminated")
}
