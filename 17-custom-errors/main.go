package main

// anything that has an Error function returning a string
// satisfies the contract of error interface
//
// type error interface {
// 	Error() string
// }

import (
	"fmt"
)

type ParseError struct {
	Message string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d \n"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}

func main() {

	bigError := ParseError{"big problem", 7,2}
	fmt.Print(bigError.Error())

}