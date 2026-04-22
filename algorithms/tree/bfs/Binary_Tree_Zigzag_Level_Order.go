/*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverse(a []int) []int{
    l,r := 0, len(a)-1
    for l < r{
        a[l], a[r] = a[r], a[l]
        l++
        r--
    }
    return a
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil{return [][]int{}}
    var ans [][]int
    queue := []*TreeNode{root}
    l_r := true
    for len(queue) > 0{
        size := len(queue)
        var curr []int
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]

            if front.Left != nil{queue = append(queue, front.Left)}
            if front.Right != nil{queue = append(queue, front.Right)}

            curr = append(curr, front.Val)
        }
        if l_r{
            ans = append(ans, curr)
        } else{
            ans = append(ans, reverse(curr))
        }

        l_r = !l_r
    }
    return ans 
}