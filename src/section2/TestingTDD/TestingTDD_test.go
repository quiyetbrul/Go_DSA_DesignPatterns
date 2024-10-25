package main

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expected := 3
	result := add(a, b)
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
