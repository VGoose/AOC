package day1

import (
	"bufio"
	"os"
	"strconv"
)

func SonarSweep(fileName string) int {
	f, err := os.Open(fileName)

	if err != nil {
		return 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// read first line
	scanner.Scan()
	prev, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0
	}

	count := 0
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0
		}

		if depth > prev {
			count++
		}
		prev = depth
	}

	return count
}

func SonarSweep2(fileName string) int {
	f, err := os.Open(fileName)

	if err != nil {
		return 0
	}
	defer f.Close()

	// scanner := bufio.NewScanner(f)
	// a, b, c := 0, 0, 0
	// scanner.Scan()
	// n, err := strconv.Atoi(scanner.Text())

	// if err != nil {
	// 	return 0
	// }

	// a = n

	// scanner.Scan()
	// n, err = strconv.Atoi(scanner.Text())

	// if err != nil {
	// 	return 0
	// }

	// a += n
	// b = n

	return 0
}
