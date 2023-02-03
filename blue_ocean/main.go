package main

import "fmt"

func main() {
	list1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list2 := []int{1}

	res := blueOcean(list1, list2)

	fmt.Println(res)
}

func blueOcean(list1, list2 []int) []int {
	temp := make(map[int]bool, len(list2))

	for _, v := range list2 {
		temp[v] = true
	}

	result := make([]int, 0, len(list1))
	for _, v := range list1 {
		if !temp[v] {
			result = append(result, v)
		}
	}

	return result
}
