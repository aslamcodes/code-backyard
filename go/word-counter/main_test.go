package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	input := bytes.NewBufferString("word1 word2")

	exp := 2

	actual := count(input, false)

	if exp != actual {
		t.Errorf("Nope")
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("word1 word2")

	exp := 1

	actual := count(input, true)

	if exp != actual {
		t.Errorf("Nope")
	}
}
