package design

import (
	"math/rand"
	"time"
)

//https://leetcode.com/explore/featured/card/top-interview-questions-easy/98/design/670/

type Solution struct {
	nums []int
	orig []int
}

func Constructor(nums []int) Solution {
	t := make([]int, len(nums))
	copy(t, nums)
	return Solution {
		nums: nums,
		orig: t,
	}
}

// Reset resets the array to its original configuration and return it.
func (this *Solution) Reset() []int {
	t := make([]int, len(this.orig))
	copy(t, this.orig)
	this.nums=t
	return t
}


//Shuffle returns a random shuffling of the array.
func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(this.nums), func(i, j int) { this.nums[i], this.nums[j] = this.nums[j], this.nums[i] })
	return this.nums
}
