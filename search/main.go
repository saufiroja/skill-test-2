package main

func main() {
	color := []string{"red", "blue", "green", "yellow", "pink", "black"}
	target := "yellow"
	res := search(color, target)
	println(res)
}

func search(color []string, target string) int {
	for _, v := range color {
		if v == target {
			return 1
		}
	}
	return -1
}
