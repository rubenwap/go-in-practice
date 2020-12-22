package main

import (
)

//MyWriter is a fake struct
type MyWriter struct {

}

func (m *MyWriter) Write([]byte) error {
	//this doesn't really implement io.Writer, even if it looks like it does
	return nil
}

func main() {
	
}
