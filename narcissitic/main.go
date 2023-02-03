package main

import (
	"math"
	"strconv"
)

func main() {
	nums := 111
	println(narcissistic(nums))
}

func narcissistic(value int) bool {
	str := strconv.Itoa(value)
	sum := 0

	for _, v := range str {
		num, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(num), float64(len(str))))
	}

	return sum == value
}
