package main

import (
	"fmt"
	"sort"
)

/*
https://www.hello-algo.com/chapter_backtracking/backtracking_algorithm/
回溯算法（backtracking algorithm）是一种通过穷举来解决问题的方法，它的核心思想是从一个初始状态出发，
暴力搜索所有可能的解决方案，当遇到正确的解则将其记录，直到找到解或者尝试了所有可能的选择都无法找到解为止。

回溯算法通常采用“深度优先搜索”来遍历解空间
*/

// 全排列问题1
// 无相等元素的情况
// 输入一个整数数组，其中不包含重复元素，返回所有可能的排列。
func permute1(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			tmp := append([]int{}, path...) // 拷贝一份，避免引用复用
			res = append(res, tmp)
			return
		}

		// 遍历所有选择
		for i := 0; i < len(nums); i++ {
			// 剪枝：不允许重复选择元素
			if used[i] {
				continue
			}

			//尝试：做出选择，更新状态
			used[i] = true
			path = append(path, nums[i])
			// 进行下一轮选择
			backtrack()
			// 回退：撤销选择，恢复到之前的状态
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// 全排列问题1
// 考虑相等元素的情况
// 输入一个整数数组，数组中可能包含重复元素，返回所有不重复的排列。
func permute2(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	// 1. 排序，使得重复元素相邻
	sort.Ints(nums)

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			tmp := append([]int{}, path...) // 拷贝一份，避免引用复用
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			// 关键：去重逻辑，这个技巧可以保证在每一层递归中，相同数字只会被选择一次，从而去掉重复解。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			//尝试：做出选择，更新状态
			used[i] = true
			path = append(path, nums[i])
			// 进行下一轮选择
			backtrack()
			// 回退：撤销选择，恢复到之前的状态
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

func main() {
	//res := permute1([]int{7, 199, 1})
	res := permute2([]int{7, 199, 7})

	for _, v := range res {
		fmt.Println(v)
	}
}
