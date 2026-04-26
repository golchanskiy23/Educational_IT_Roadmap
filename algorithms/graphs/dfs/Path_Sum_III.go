/*
targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, curr, ts int, m map[int]int) int{
    if root == nil{return 0}
    cnt := 0
    curr += root.Val
    cnt += m[curr-ts]
    m[curr]++
    l := dfs(root.Left, curr, ts, m)
    r := dfs(root.Right, curr, ts, m)
    m[curr]--
    if m[curr] == 0{delete(m, curr)}
    return l+r+cnt
}

func pathSum(root *TreeNode, ts int) int {
    m := map[int]int{0: 1}
    return dfs(root, 0, ts, m)
}