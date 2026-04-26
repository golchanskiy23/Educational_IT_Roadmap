/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func dfs(curr *Node, visited map[*Node]*Node) *Node{
    if cloned , ok := visited[curr]; ok{
        return cloned
    }

    node := &Node{Val: curr.Val}
    visited[curr] = node
    for _, n := range curr.Neighbors{
        node.Neighbors = append(node.Neighbors, dfs(n, visited))
    }

    return node
}

func cloneGraph(node *Node) *Node {
    if node == nil{return nil}
    m := make(map[*Node]*Node)
    dfs(node, m) 
    return m[node]
}