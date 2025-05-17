package main

import (
	"flag"
	"fmt"
	"todo"
)

func main() {
	var vault = flag.String("v", "home", "The vault name")

	flag.Parse()

	var list = todo.List{}

	err := list.Load(*vault)

	if err != nil {
		// How I'm going to be sure that it's only going to be file not found error?
		fmt.Printf("Creating vault on %s", *vault)
		list.Save(*vault)
	}

	for _, todo := range flag.Args() {
		list.Add(todo)
	}

	for _, todo := range list {
		fmt.Println(todo.Title)
	}

	list.Save(*vault)
}
