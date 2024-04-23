package main

import "fmt"

// twoSum finds two numbers such that they add up to a specific target
// and returns the indices of the two numbers.
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Create a map to store numbers and their indices

	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // Output: [0 1]
}
