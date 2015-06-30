package main

import (
	"fmt"
	"time"
)


func main() {
	go func(i int) {
		fmt.Println(i)
	}(3)

	time.Sleep(1)
	fmt.Println("Terminated")	
}
