package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, lines bool) int {
	wc := 0

	sc := bufio.NewScanner(r)

	if !lines {
		sc.Split(bufio.ScanWords)
	}

	for sc.Scan() {
		wc++
	}

	return wc
}

func main() {
	var line = flag.Bool("lines", false, "Counts the lines instead")

	flag.Parse()

	fmt.Println(count(os.Stdin, *line))
}
