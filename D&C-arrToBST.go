package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func add(num int, tree *TreeNode) {
	if tree.Val > num {
		if tree.Left == nil {
			tree.Left = &TreeNode{Val: num, Left: nil, Right: nil}
			return
		}
		add(num, tree.Left)
	} else if tree.Val < num {
		if tree.Right == nil {
			tree.Right = &TreeNode{Val: num, Left: nil, Right: nil}
			return
		}
		add(num, tree.Right)
	}
}

func recursiveCall(nums []int, tree *TreeNode) {
	l := len(nums)
	if l > 2 {
		left := nums[0 : (l/2)+1]
		right := nums[l/2+1:]
		recursiveCall(left, tree)
		recursiveCall(right, tree)
	}

	if l == 1 {
		add(nums[0], tree)
		return
	}
	add(nums[1], tree)
	add(nums[0], tree)
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	middleValue := nums[l/2]
	tree := &TreeNode{Val: middleValue, Left: nil, Right: nil}

	if l > 2 {
		left := nums[0:(l / 2)]
		right := nums[(l/2)+1:]
		recursiveCall(left, tree)
		recursiveCall(right, tree)
	} else if l == 2 {
		tree.Left = &TreeNode{Val: nums[0], Left: nil, Right: nil}
	}

	return tree
}
