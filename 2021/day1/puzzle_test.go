package day1

import "testing"

func TestPuzzle(t *testing.T) {
	n := SonarSweep("input.txt")
	if n != 1715 {
		t.Fatalf("Expected 1715 but got %d", n)
	}
}

func TestPuzzle2(t *testing.T) {
	n := SonarSweep2("input.txt")
	if n != 1739 {
		t.Fatalf("Expected 1739 but got %d", n)
	}
}
