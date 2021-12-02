package day4

import (
	"bufio"
	"os"
	"strings"
)

func passportScanner(fileName string) int {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return 0
	}

	count := 0
	codes := map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		if t == "" {
			// if not empty
			// tally & reset
			if len(codes) != 0 {
				if len(codes) >= 7 {

				}
				codes = map[string]bool{}
			}
			continue
		}
		groups := strings.Split(t, " ")

		for _, code := range groups {
			codes[code[:3]] = true
		}
	}
	return count
}
