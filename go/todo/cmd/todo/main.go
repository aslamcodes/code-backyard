package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

func main() {
	var vault = flag.String("v", "", "The vault name")

	flag.Parse()

	var list = todo.List{}

	list.Add("Hello")
	list.Save("home")

	if *vault == "" {
		fmt.Print("Please provide a vault name")
		os.Exit(-1)
	}

	list.Load(*vault)
}
