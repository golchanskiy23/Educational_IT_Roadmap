/*
A company has n employees with a unique ID for each employee from 0 to n - 1. The head of the company is the one with headID.

Each employee has one direct manager given in the manager array where manager[i] is the direct manager of the i-th employee, manager[headID] = -1. Also, it is guaranteed that the subordination relationships have a tree structure.

The head of the company wants to inform all the company employees of an urgent piece of news. He will inform his direct subordinates, and they will inform their subordinates, and so on until all employees know about the urgent news.
*/

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func dfs(adj [][]int, inform []int, id int) int{
    if inform[id] == 0{return 0}
    curr := 0
    for _, node := range adj[id]{
        curr = max(curr, dfs(adj, inform, node))
    }

    return curr + inform[id]
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    adj := make([][]int, n)
    for i := 0; i < n; i++ {adj[i] = make([]int, 0)}

    for i := 0; i < n; i++{
        if manager[i] != -1{
            adj[manager[i]] = append(adj[manager[i]], i)
        }
    }

    return dfs(adj, informTime, headID)
}