/*
Given the root of a binary tree, return all duplicate subtrees.

For each kind of duplicate subtrees, you only need to return the root node of any one of them.

Two trees are duplicate if they have the same structure with the same node values.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    m := make(map[string]int)
    ans := []*TreeNode{}

    var dfs func(*TreeNode) string
    dfs = func(root *TreeNode) string{
        if root == nil {return "#"}
        serial := fmt.Sprintf("%d, %s, %s", root.Val, dfs(root.Left), dfs(root.Right))
        m[serial]++
        if m[serial] == 2{
            ans = append(ans, root)
        }
        return serial
    }

    dfs(root)
    for k,v := range m{
        fmt.Println(k,v)
    }
    return ans
}