package day2

import "testing"

func TestPuzzle(t *testing.T) {
	n := Dive("input.txt")

	if n != 5 {
		t.Fatalf("Expected 5 but got %d", n)
	}
}

func TestPuzzle2(t *testing.T) {
	n := Dive2("input.txt")

	if n != 1842742223 {
		t.Fatalf("Expected 1842742223 but got %d", n)
	}
}
