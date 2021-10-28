// https://leetcode.com/problems/binary-tree-inorder-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    cur := root
    visited := []*TreeNode {  }

    for cur != nil || len(visited) > 0 {
        // visit record all the visits till the left most node
        for cur != nil {
            visited = append(visited, cur)
            cur = cur.Left
        }
        
        // pop the latest visit
        cur = visited[len(visited)-1]
        visited = visited[:len(visited)-1]
        
        // appaend the value visited
        res = append(res, cur.Val)
        
        // move to right
        cur = cur.Right
    }
    
    return res
}


