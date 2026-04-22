/*
Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    var matrix [][]int
    queue := []*TreeNode{root}
    for len(queue) > 0{
        size := len(queue)
        var tmp []int
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            
            if front.Left != nil{queue = append(queue, front.Left)}
            if front.Right != nil{queue = append(queue, front.Right)}

            tmp = append(tmp, front.Val)
        }
        matrix = append(matrix, tmp)
    }

    ans := make([]float64,0, len(matrix))
    for i := 0; i < len(matrix); i++{
        sum := 0.0 // float64
        for j := 0; j < len(matrix[i]); j++{
            sum += float64(matrix[i][j])
        }
        ans = append(ans, (sum / float64(len(matrix[i]))))
    }

    return ans
}