package day1

import "testing"

func TestPuzzle(t *testing.T) {
	n := SonarSweep("input.txt")
	if n != 1715 {
		t.Fatalf("Expected 1715 but got %d", n)
	}
}
