package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
"firstName" : "Jean",
"lastName" : "Bartik"
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m := f.(map[string]interface{})
	fmt.Print(m["firstName"])
}
