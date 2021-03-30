package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("1 2 33 4 5"))
}

func HighAndLow(in string) string {
	a := strings.Split(in, " ")
	b := make([]int, len(a))
	// Create []int
	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}
	max, min := findMinAndMax(b)
	return fmt.Sprintf("%d %d", max, min)
}

func findMinAndMax(array []int) (min int, max int) {
	min = array[0]
	max = array[0]
	for _, value := range array {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max, min
}
