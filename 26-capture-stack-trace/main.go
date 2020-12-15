package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false) // bool determines whether we print trace for goroutines too
	fmt.Printf("Trace: \n %s\n", buf)
}