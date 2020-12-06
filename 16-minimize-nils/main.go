package main

import (
	"fmt"
	"os"
	"errors"
	"strings"
)

func concat (parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}
	return strings.Join(parts, " "), nil
}

func main() {
	args := os.Args[1:]
	result, _ := concat(args...)
	fmt.Printf("Concatenated strings: '%s'\n", result)
	

}