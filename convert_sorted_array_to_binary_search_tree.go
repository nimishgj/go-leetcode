package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    lengthOfNums := len(nums)
    if lengthOfNums < 1 {
        return nil
    }
    midIndex := lengthOfNums / 2
    rootNode := &TreeNode{Val: nums[midIndex]}
    rootNode.Left = sortedArrayToBST(nums[:midIndex]);
    rootNode.Right = sortedArrayToBST(nums[midIndex+1:]);
    return rootNode;
}
