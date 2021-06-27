package main

import "fmt"

func slideWindow(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return nil
	}
	ret := []int{}
	que := []int{}
	for i := 0; i < len(nums); i++ {
		for len(que) > 0 && nums[i] >= nums[que[len(que)-1]] {
			que = que[:len(que)-1]
		}
		que = append(que, i)
		if que[0] < i-k+1 {
			que = que[1:]
		}
		if i >= k-1 {
			ret = append(ret, nums[que[0]])
		}
	}
	return ret
}

func main() {
	fmt.Println(slideWindow([]int{5, 4, 3, 2, 6, 7, 5, 8, 0, 9}, 3))
}