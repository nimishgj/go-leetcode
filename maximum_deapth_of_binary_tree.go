package main

func maxDepth(root *TreeNode) int {
      if root == nil {
            return 0
        }
      
      return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

  // This is DFS Problem. Using recursion we find the max deapth.
