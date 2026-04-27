package main

import "fmt"

func main() {
	nums := []int{2, 4, 3, 6}

	max := nums[0]
	second := nums[0]

	for _, v := range nums {
		if v > max {
			second = max
			max = v
		} else if v > max && v != max {
			second = v
		}

	}
	fmt.Println(second)
}
