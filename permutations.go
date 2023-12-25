package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	var result [][]int
	generatePermutations(nums, 0, &result)
	return result
}

func generatePermutations(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		// Append a copy of the current permutation to the result
		*result = append(*result, append([]int{}, nums...))
		return
	}

	for i := start; i < len(nums); i++ {
		// Swap the current element with the element at index 'start'
		nums[start], nums[i] = nums[i], nums[start]

		// Recursively generate permutations for the remaining elements
		generatePermutations(nums, start+1, result)

		// Backtrack: restore the original order before the next iteration
		nums[start], nums[i] = nums[i], nums[start]
	}
}
