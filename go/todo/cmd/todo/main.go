package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"todo"
)

func main() {
	var vault = flag.String("v", "home", "The vault name")
	var listTasks = flag.Bool("l", false, "List the tasks")
	var completeTask = flag.Int("c", 0, "The item to complete")
	var showAll = flag.Bool("all", false, "Shows all the tasks")

	var profiling = flag.Bool("p", false, "Profiling")

	flag.Parse()

	if *profiling {
		defer trackTime(time.Now(), "main")
	}

	var list = todo.List{}

	todo.Init(*vault)

	if err := list.Load(*vault); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(flag.Args()) > 0 {
		list.Add(strings.Join(flag.Args(), " "))
	}

	switch {
	case *listTasks:
		for _, item := range list {
			if !item.Done || *showAll {
				fmt.Println(item.Title)
			}
		}
	case *completeTask > 0:
		list.Complete(*completeTask)
	}

	err := list.Save(*vault)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
