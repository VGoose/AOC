package day3

import (
	"testing"
)

func TestDay3(t *testing.T) {
	n := countTrees("./input.txt")

	if n != 200 {
		t.Fatalf("Expected 200, got %d.", n)
	}
}

func TestDay3_2(t *testing.T) {
	n := countTrees2("./input.txt")

	if n != 3737923200 {
		t.Fatalf("Expected 3737923200, but got %d", n)
	}
}
