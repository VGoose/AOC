package day2

import "testing"

func TestPuzzle(t *testing.T) {
	n := Dive("input.txt")

	if n != 5 {
		t.Fatalf("Expected 5 but got %d", n)
	}
}
