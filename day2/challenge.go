package day2

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// parse "1-3 a: abcde" && validate letter occurs within a given range in the password
func Challenge(l []string) int {
	count := 0
	for _, s := range l {
		low, high, letter, password, ok := parseArgs(s)
		if !ok {
			continue
		}

		c := 0
		for _, r := range password {
			if r == letter {
				c++
			}
		}
		if c >= low && c <= high {
			count++
		}
	}
	return count
}

// parse "1-3 a: abcde" && validate letter occurs at _one_ of the given positions in the password
func Challenge2(l []string) int {
	count := 0
	for _, s := range l {
		i, j, letter, password, ok := parseArgs(s)

		if !ok {
			continue
		}

		p := []rune(password)
		if p[i-1] == letter && p[j-1] == letter {
			continue
		}
		if p[i-1] == letter || p[j-1] == letter {
			count++
		}
	}
	return count
}
func parseArgs(s string) (int, int, rune, string, bool) {
	args := strings.Split(s, " ")
	lowHigh := strings.Split(args[0], "-")
	low, err := strconv.Atoi(lowHigh[0])
	high, err := strconv.Atoi(lowHigh[1])
	ok := true

	if err != nil {
		ok = false
	}

	letter, _ := utf8.DecodeRuneInString(args[1])
	return low, high, letter, args[2], ok
}
