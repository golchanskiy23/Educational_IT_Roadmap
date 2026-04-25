/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
*/

func dfs(adj [][]int, visited []int, curr int) bool{
    if visited[curr] == 1{return false}
    if visited[curr] == 2{return true}

    visited[curr] = 1
    for _, node := range adj[curr]{
        if !dfs(adj, visited, node){
            return false
        }
    }
    visited[curr] = 2

    return true
}

func canFinish(n int, pre [][]int) bool {
    // топологическая сортировка
    // if graph[pre[i][0]] = pre[i][1]{false}
    // но не проверяет уже связные узлы
    // если не посещён - белый - 0
    // посещён, но не обработан до конца - grey - 1
    // посещён и обработан - black - 2

    adj := make([][]int, n)
    for i := 0;i < n; i++{adj[i] = make([]int, 0)}

    for i := 0; i< len(pre); i++{
        visited := make([]int, n)
        adj[pre[i][1]] = append(adj[pre[i][1]], pre[i][0])
        if !dfs(adj, visited, pre[i][1]){
            return false
        }
    }
    
    return true
}