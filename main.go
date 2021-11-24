package main

import "fmt"

func main() {
	fmt.Println(day1(day1_input, 2020))
}

// find 2 entries that add up to 2020 & return their product
func day1(l []int, t int) int {
	compl := map[int]bool{}
	for _, v := range l {
		if compl[t-v] {
			return v * (t - v)
		}
		compl[v] = true
	}
	return -1
}
