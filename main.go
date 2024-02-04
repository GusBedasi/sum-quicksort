package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 3, 2, 1, 4, 10, 20, 30, 500, 5000, 14560, 2313, 1111, 7}
	fmt.Println(sum(nums))
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	pivot, queue := nums[0], nums[1:]

	return pivot + sum(queue)
}
