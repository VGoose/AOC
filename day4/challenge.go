package day4

import (
	"bufio"
	"fmt"
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
			continue
		}
		groups := strings.Split(t, " ")

		for _, code := range groups {
			codes[code[:3]] = true
			fmt.Printf("%v", codes)
			fmt.Print("\n")
		}
	}
	return count
}
