/*
There are n computers numbered from 0 to n - 1 connected by ethernet cables connections forming a network where connections[i] = [ai, bi] represents a connection between computers ai and bi. Any computer can reach any other computer directly or indirectly through the network.

You are given an initial computer network connections. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected.

Return the minimum number of times you need to do this in order to make all the computers connected. If it is not possible, return -1.
*/

func dfs(adj [][]int, visited []bool, curr int){
    if visited[curr]{return}
    
    visited[curr] = true
    for _, node := range adj[curr]{
        dfs(adj, visited, node)
    }
}

func makeConnected(n int, connections [][]int) int {
    adj := make([][]int, n)
    for i := 0; i < n; i++{
        adj[i] = make([]int, 0)
    }

    for i := 0; i < len(connections); i++{
        adj[connections[i][0]] = append(adj[connections[i][0]], connections[i][1]) 
        adj[connections[i][1]] = append(adj[connections[i][1]], connections[i][0]) 
    }

    visited := make([]bool, n)
    components := 0
    for i := 0; i < n; i++{
        if !visited[i]{
            dfs(adj, visited, i)
            components++
        }
    }
    if len(connections) < n-1 {return -1} 

    return components-1
}