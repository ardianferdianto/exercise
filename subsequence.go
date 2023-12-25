package main

import (
	"fmt"
)

func main() {
	c := []int{1, 3, 5, 4, 7, 8, 9, 10}
	fmt.Printf("result %+v\n", findLengthOfLCIS(c))

}

func findLengthOfLCIS(nums []int) int {
	count := 1
	max := 0
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(count, nums[i])
	// 	if nums[i] < nums[i+1] {
	// 		count++
	// 	} else {
	// 		if max < count {
	// 			max = count
	// 		}
	// 		count = 1
	// 	}
	// }

	for i, num := range nums {
		fmt.Println(count, num)
		if i+1 < len(nums) && num < nums[i+1] {
			count++
		} else {
			if max < count {
				max = count
			}
			count = 1
		}
	}
	return max
}
