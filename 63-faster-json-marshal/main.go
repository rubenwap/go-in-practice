package main

import (
	"fmt"
	"os/user"
)

func main() {

	jh := new(codec.JaonHandle)
	u := &user.User{
		Name:  "Inigo Montoya",
		Email: "inigo@example.com",
	}

	var out []byte
	err := codec.NewEncoderBytes(&out, jh).Encode(&u)
	if err != nil {

	}

	fmt.Println(string(out))
}
