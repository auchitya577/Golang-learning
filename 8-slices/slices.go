package main

import (
	"fmt"
	
)

func main() {
	//uninitialized slicee is nil

	// var nums []int
	// fmt.Println(nums==nil)
	//  fmt.Println(len(nums))

	// var nums = make([]int,0,5)
     // IT IS LENGTH,CAPACITY INT K BJU M
	//cap-> MAX NUMBERS OF ELEMENTS WHICH CAN FIX IN THIS SLICE
     //  fmt.Println(cap(nums))
	// fmt.Println(nums)
     
	// nums:=[]int{}
	// nums=append(nums, 1)
	// nums=append(nums, 5)
	// nums=append(nums, 9)

	// nums[0]=3
	// nums[1]=5
	// nums[2]=54
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	//COPY FUNCTION
	// var nums = make([]int,0,5)
	// nums=append(nums, 2)
    // var nums2=make([]int,len(nums))
	// nums2[0]=2
	// nums2 = append(nums2, 20)
	
	// //copy function

	// // copy(nums2,nums)
	// fmt.Println(nums,nums2)

	//SLICE OPERATOR

	// var nums=[]int{1,2,3}
	// fmt.Println(nums[0:3])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	//SLICE PACKAGE(used for comparison elements gives true or false)

	// var nums1=[]string{"auchit"}
	// var nums2=[]string{"auchity"}
    // fmt.Println(slices.Equal(nums1,nums2))

	var nums=[][]int{{1,2,3},{4,5,6,}}
	fmt.Println(nums)
}
