package main

import (
	"fmt"
	"runtime"
//	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	done := make(chan bool, 1)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("go routine: ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<- done
		fmt.Println("main routine", i)
		
	}
}
