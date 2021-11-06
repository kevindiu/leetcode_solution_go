// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    ress := make([]int, 0, k) 
    dfs(root, &ress, k)
    return ress[k-1]
}

func dfs(n *TreeNode, res *[]int, k int) {
    if n == nil {
        return
    }
    
    dfs(n.Left, res, k)
    
    *res = append(*res, n.Val)
    if len(*res) == k {
        return
    }
    
    dfs(n.Right, res, k)
}
