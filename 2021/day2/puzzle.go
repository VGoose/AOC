package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Dive(fileName string) int {
	f, err := os.Open(fileName)
	if err != nil {
		return 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	x, y := 0, 0
	for scanner.Scan() {
		t := scanner.Text()
		ins := strings.Split(t, " ")
		units, err := strconv.Atoi(ins[1])
		if err != nil {
			return 0
		}

		switch ins[0] {
		case "forward":
			x += units
		case "up":
			y -= units
		case "down":
			y += units
		}
	}
	return x * y
}
