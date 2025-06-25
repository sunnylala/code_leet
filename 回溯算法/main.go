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

// 全排列问题2
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

//func main() {
//	//res := permute1([]int{7, 199, 1})
//	res := permute2([]int{7, 199, 7})
//
//	for _, v := range res {
//		fmt.Println(v)
//	}
//}

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

/*
前序遍历：例题三
在二叉树中搜索所有值为7的节点，请返回根节点到这些节点的路径，
并要求路径中不包含值为3的节点。
*/
func preOrderIII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	// 剪枝
	if root == nil || root.Val == 3 {
		return
	}
	// 尝试
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderIII(root.Left, res, path)
	preOrderIII(root.Right, res, path)
	// 回退
	*path = (*path)[:len(*path)-1]
}

func preOrderSearchPaths(root *TreeNode) [][]*TreeNode {
	var res [][]*TreeNode
	var path []*TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		// 剪枝
		if node == nil || node.Val == 3 {
			return
		}
		// 尝试
		path = append(path, node)
		if node.Val == 7 {
			tmp := make([]*TreeNode, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		dfs(node.Left)
		dfs(node.Right)
		// 回退
		path = path[:len(path)-1]
	}

	dfs(root)
	return res
}

//
//func main() {
//	// 构造示例二叉树
//	/*
//	        1
//	      /   \
//	     2     3
//	    / \   / \
//	   7   4 5   7
//	          \
//	           3 (无效路径，包含3)
//	*/
//	root := &TreeNode{Val: 1}
//	root.Left = &TreeNode{Val: 2}
//	root.Right = &TreeNode{Val: 3}
//	root.Left.Left = &TreeNode{Val: 7}
//	root.Left.Right = &TreeNode{Val: 4}
//	root.Right.Left = &TreeNode{Val: 5}
//	root.Right.Right = &TreeNode{Val: 7}
//	root.Right.Left.Right = &TreeNode{Val: 3}
//
//	// 搜索路径
//	paths := preOrderSearchPaths(root)
//
//	// 打印结果
//	fmt.Println("符合条件的路径：")
//	for _, path := range paths {
//		for i, node := range path {
//			if i > 0 {
//				fmt.Print(" -> ")
//			}
//			fmt.Print(node.Val)
//		}
//		fmt.Println()
//	}
//}

//给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，使得组合中的元素和等于 target 。
//给定数组无重复元素，每个元素可以被选取多次。请以列表形式返回这些组合，列表中不应包含重复组合。
//输入集合中的元素可以被无限次重复选取。
/* 回溯算法：子集和 I */
/* 求解子集和 I（包含重复子集） */
func subsetSumINaive(nums []int, target int) [][]int {
	path := make([]int, 0)  // 状态（子集）
	res := make([][]int, 0) // 结果列表（子集列表）

	var dfs func(start, total int)
	dfs = func(start, total int) {
		// 子集和等于 target 时，记录解
		if target == total {
			tmp := append([]int{}, path...)
			res = append(res, tmp)
			return
		}
		if total > target {
			return
		}

		// 遍历所有选择
		for i := start; i < len(nums); i++ {
			// 剪枝：若子集和超过 target ，则跳过该选择
			if total+nums[i] > target {
				continue
			}
			// 选择 nums[i]
			path = append(path, nums[i])
			// 进行下一轮选择
			dfs(i, total+nums[i])
			// 回溯
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return res
}

// 给定一个正整数数组 nums 和一个目标正整数 target ，请找出所有可能的组合，使得组合中的元素和等于 target 。
// 给定数组可能包含重复元素，每个元素只可被选择一次。请以列表形式返回这些组合，列表中不应包含重复组合。
func subsetSumINaive2(nums []int, target int) [][]int {
	sort.Ints(nums) // 排序以便剪枝与去重
	var res [][]int
	var path []int

	var dfs func(start int, total int)
	dfs = func(start int, total int) {
		if total == target {
			tmp := append([]int{}, path...)
			res = append(res, tmp)
			return
		}
		if total > target {
			return
		}

		for i := start; i < len(nums); i++ {
			// 跳过同层重复元素（避免重复组合）
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			// 选择 nums[i]
			path = append(path, nums[i])
			// 因为每个元素只能用一次，递归时从 i+1 开始
			dfs(i+1, total+nums[i])
			// 回溯
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return res
}

func main() {
	nums := []int{2, 3, 4, 7}
	target := 7
	combinations := subsetSumINaive2(nums, target)

	for _, comb := range combinations {
		fmt.Println(comb)
	}
}
