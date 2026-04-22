/*
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func dfs(root *TreeNode, ans *int) int{
    if root == nil{return 0}
    l_height, r_height := dfs(root.Left, ans), dfs(root.Right, ans)

    *ans = max(*ans, l_height+r_height)
    
    return max(l_height, r_height)+1
}

func diameterOfBinaryTree(root *TreeNode) int {
    var ans int = 0
    dfs(root, &ans)
    return ans
}