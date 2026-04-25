/*
You are given an integer n, which indicates that there are n courses labeled from 1 to n. You are also given a 2D integer array relations where relations[j] = [prevCoursej, nextCoursej] denotes that course prevCoursej has to be completed before course nextCoursej (prerequisite relationship). Furthermore, you are given a 0-indexed integer array time where time[i] denotes how many months it takes to complete the (i+1)th course.

You must find the minimum number of months needed to complete all the courses following these rules:

You may start taking a course at any time if the prerequisites are met.
Any number of courses can be taken at the same time.
Return the minimum number of months needed to complete all the courses.

Note: The test cases are generated such that it is possible to complete every course (i.e., the graph is a directed acyclic graph).
*/

func max(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func dfs(adj [][]int, time, dist []int, curr int) int{
    if dist[curr] != 0 {return dist[curr]}

    currMax := 0
    for _, node := range adj[curr]{
        currMax = max(currMax, dfs(adj, time, dist, node))
    }
    dist[curr] = currMax+time[curr-1]
    return dist[curr]
}

func minimumTime(n int, relations [][]int, time []int) int {
    dist := make([]int, n+1)

    adj := make([][]int, n+1)
    for i := 0; i <= n; i++{
        adj[i] = make([]int, 0)
    }

    for i := 0; i < len(relations); i++{
        adj[relations[i][1]] = append(adj[relations[i][1]], relations[i][0])
    }

    ans := 0
    for i := 1; i <= n; i++{
        ans = max(ans, dfs(adj, time, dist, i))
    }
    
    return ans
}