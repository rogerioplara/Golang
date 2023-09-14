package main

import "fmt"

func main() {
	nums := []int{2, 5, 5, 11}
	target := 10

	resultado := twoSum(nums, target)

	fmt.Println(resultado)
}

func twoSum(nums []int, target int) []int {
	var s []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				s = append(s, i, j)
				return []int{i, j}
			}
		}
	}
	return []int{}
}
