package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside goroutine")
	go func() {
		fmt.Println("Inside goroutine")
	}()
	fmt.Println("Outside again")
	runtime.Gosched()
}