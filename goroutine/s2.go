package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
//	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("go routine: ", i)
		}()
	}
	fmt.Scanln()
}
