package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	list.Add(strings.Join(flag.Args(), " "))

	// Show Items
	for _, todo := range list {
		fmt.Println(todo.Title)
	}

	err := list.Save(*vault)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
