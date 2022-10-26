package main

import (
	"strings"
	"testing"
)

func TestLenString(t *testing.T) {
	input := "Hello Worls"
	expected := 11
	result := len(input)

	if result != expected {
		t.Errorf("Incorrect result. Except %d, got %d", expected, result)
	}
}
func TestContainString(t *testing.T) {
	input := "Hello Guys"
	word := "Guy"
	expected := true
	result := strings.Contains(input, word)

	if result != expected {
		t.Errorf("Incorrect result.")
	}
}

func TestSearchString(t *testing.T) {
	input := "Hello"
	number := 4
	expected := 111
	result := input[number]

	if int(result) != expected {
		t.Errorf("Incorrect result. Except %d, got %d", expected, result)
	}
}
