package main

import "fmt"

func main() {
	nums := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	parityOutlier(nums)
}

func parityOutlier(arr []int) {
	even := 0
	odd := 0
	for _, v := range arr {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	switch {
	case even == 1:
		for _, v := range arr {
			if v%2 == 0 {
				fmt.Println(v)
			}
		}
	case odd == 1:
		for _, v := range arr {
			if v%2 != 0 {
				fmt.Println(v)
			}
		}
	default:
		fmt.Println(false)
	}

}
