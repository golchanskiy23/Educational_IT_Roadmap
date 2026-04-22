/*
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 2 ребенка - заменяем на максимум и обнуляем(правого/левого)
// 1 ребенок - заменяем и обнуляем
// нет детей - зануляем узел

func findMin(root *TreeNode) *TreeNode{
    if root == nil{return nil}
    for root.Left != nil{
        root = root.Left
    }

    return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil{return nil}

    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val{
        root.Right = deleteNode(root.Right, key)
    } else{
        if root.Left == nil && root.Right == nil{
            return nil
        } else if root.Left != nil && root.Right != nil{
            node := findMin(root.Right)
            root.Val = node.Val
            root.Right = deleteNode(root.Right, node.Val)
        } else{
            if root.Left != nil{
                root = root.Left
            } else{
                root = root.Right
            }
        }
    }

    return root
}