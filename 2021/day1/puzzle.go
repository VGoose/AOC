package day1

import (
	"bufio"
	"math"
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

	scanner := bufio.NewScanner(f)
	window := []int{}
	prev := math.MaxInt32
	count := 0

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0
		}

		window = append(window, n)

		if len(window) != 3 {
			continue
		}

		sum := window[0] + window[1] + window[2]
		if sum > prev {
			count++
		}
		prev = sum
		window = window[1:]
	}

	return count
}
