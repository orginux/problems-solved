package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	// Output: [5,6,7,1,2,3,4]
}

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		previous := nums[len(nums)-1]
		fmt.Printf("----- #%d \n", i)
		for j, _ := range nums {
			nums[j], previous = previous, nums[j]
			fmt.Printf("\t replace %d to %d: %v\n", previous, nums[j], nums)
		}
	}

	fmt.Println(nums)
}
