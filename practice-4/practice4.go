package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	target := 5

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {

				fmt.Println(nums[i], nums[j])
			}
		}
	}
}
