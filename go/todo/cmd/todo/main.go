package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

func main() {
	var vault = flag.String("v", "home", "The vault name")
	flag.Parse()

	var list = todo.List{}

	todo.Init(*vault)

	if err := list.Load(*vault); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, todo := range flag.Args() {
		list.Add(todo)
	}

	for _, todo := range list {
		fmt.Println(todo.Title)
	}

	list.Save(*vault)
}
