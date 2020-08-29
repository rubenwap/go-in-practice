package main

import (
	"fmt"
	"os"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print Hello World"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
		cli.BoolFlag{
			Name: "spanish, s",
			Usage: "Use spanish language",
		},
	}
	app.Action = func(c *cli.Context) error {
		name:= c.GlobalString("name")
		if c.Bool("spanish") {
			fmt.Printf("Hola %s\n", name)
		} else {
		fmt.Printf("Hello %s\n", name)
		}
		return nil
	}
	app.Run(os.Args)
}