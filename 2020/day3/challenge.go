package day3

import (
	"bufio"
	"os"
)

func countTrees(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		return 0
	}

	scanner := bufio.NewScanner(file)
	pos := 0
	trees := 0

	scanner.Scan()
	for scanner.Scan() {
		pos += 3
		i := pos % 31

		b := scanner.Bytes()
		if len(b) > 0 && b[i] == 35 {
			trees++
		}
	}
	return trees
}

func countTrees2(fileName string) int {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return 0
	}

	routes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	product := 0
	for j, route := range routes {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)

		pos := 0
		trees := 0
		scanner.Scan()
		for scanner.Scan() {
			// travel right & down
			pos += route[0]
			for j := 1; j < route[1]; j++ {
				scanner.Scan()
			}
			i := pos % 31

			b := scanner.Bytes()
			if len(b) > 0 && b[i] == 35 {
				trees++
			}
		}
		if j == 0 {
			product = trees
		} else {
			product *= trees
		}
	}

	return product
}
