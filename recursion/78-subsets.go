package recursion

import (
	"fmt"
)

//https://leetcode.com/problems/subsets/

func Subsets(nums []int) {
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
	return subsetBfs([][]int{}, []int{}, nums, 0)
}

func subsetBfs(solution [][]int, current, nums []int, i int) [][]int {
	if i == len(nums) { //Break out of recursive calls
		x := make([]int, len(current))
		copy(x, current)
		return append(solution, x)
	}
	//Generate all permutations that do not include the current number
	without := subsetBfs(solution, current, nums, i+1)

	//Generate all permutations that include the current number, which is appended to the variable `current`
	with := subsetBfs(solution, append(current, nums[i]), nums, i+1)
	return append(without, with...)
}

//Todo: subsetDfg (using backtracking)
//https://leetcode.com/problems/combination-sum/discuss/16502/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning)
func subsetDfs(solution [][]int, current, nums []int, i int) [][]int {
	if i == len(nums) { //Break out of recursive calls
		x := make([]int, len(current))
		copy(x, current)
		return append(solution, x)
	}

	for _ = range nums[i:] {

	}

	return solution
}