package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	wc := 0

	sc := bufio.NewScanner(r)

	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		wc++
	}

	return wc
}

func main() {
	fmt.Println(count(os.Stdin))
}
