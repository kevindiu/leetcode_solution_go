// https://leetcode.com/problems/symmetric-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return root == nil || isMirror(root.Left, root.Right) 
}

func isMirror(n1, n2 *TreeNode) bool {
    if n1 == nil || n2 == nil { 
        return n1 == n2
    }
    if n1.Val != n2.Val {
        return false
    }
    return isMirror(n1.Left, n2.Right) && isMirror(n1.Right, n2.Left)
}
