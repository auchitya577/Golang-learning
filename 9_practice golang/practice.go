package main

import "fmt"

func main() {
	//	NORMAL PRINT
	// nums := []int{1, 2, 3, 4, 5}
	// for i, v := range nums {

	// 	fmt.Println(i,v)
	// }

	//ADD NUMBERS THROUGH APPEND
	//  nums:=[]int{}
	// nums = append(nums, 20)
	// nums = append(nums, 30)
	// nums = append(nums, 50)
	// fmt.Println(nums)

	//PRINT SOME SLECETED ELEMENTS

	// nums:=[]int{20,30,40,50}
	// newslice:=nums[1:4]
	// fmt.Println(newslice)

	//SUM OF even  ELEMENTS

	nums := []int{1, 2, 3, 4, 5}

	for _, v := range nums {
		if v%2 == 0 {
			fmt.Println(v)
		}
	}
}
