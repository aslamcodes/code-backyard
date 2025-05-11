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

func countBytes(r io.Reader) int {
	sc := bufio.NewScanner(r)

	bytes := 0

	for sc.Scan() {
		bytes += len(sc.Bytes())
	}

	return bytes
}

func main() {
	// Not clearly the best way to do this, or even its is wrong
	var line = flag.Bool("l", false, "Counts the lines instead")
	var bytes = flag.Bool("b", false, "Counts bytes")

	flag.Parse()

	if *line || !*bytes {
		fmt.Println(count(os.Stdin, *line))
	} else {
		fmt.Println(countBytes(os.Stdin))
	}
}
