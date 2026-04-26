/*
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1 and return them in any order.

The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).
*/

func dfs(graph [][]int, ans *[][]int, curr *[]int, start,n int){
    *curr = append(*curr, start)
    if start == n{
        // Копируем путь, иначе все пути будут ссылаться на один срез
        tmp := make([]int, len(*curr))
        copy(tmp, *curr)
        *ans = append(*ans, tmp)
        return
    }

    for _, nei := range graph[start]{
        dfs(graph, ans, curr, nei, n)
        *curr = (*curr)[:(len(*curr)-1)]        
    }
}

func allPathsSourceTarget(graph [][]int) [][]int {
    var ans [][]int
    n := len(graph)
    var tmp []int
    dfs(graph, &ans, &tmp, 0, n-1)
    return ans
}