package main

import (
	"bytes"
	"testing"
)

func testCount(t *testing.T) {

	input := bytes.NewBufferString("word1 word2")

	exp := 2

	actual := count(input)

	if exp != actual {
		t.Errorf("Nope")
	}

}
