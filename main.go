package main

import (
	"fmt"
	"sort"
)

func Solve(nums []int) int {
	sort.Slice(nums, func (i, j int) bool {
		return nums[i] < nums[j]
	})
	count := 1
	for ; count-1 < len(nums); count++ {
		if nums[count-1] != count {
			return count
		}
	}
	return count
}

func main() {
	fmt.Println(Solve([]int{1, 2, 3, 4, 5}))
	fmt.Println(Solve([]int{4, 2, 6}))
	fmt.Println(Solve([]int{2, 1, 3, 5}))
}
