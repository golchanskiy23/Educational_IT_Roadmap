/*
You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.
*/

func getNeighbours(s string) []string{
    var ans []string
    
    for i := 0; i < len(s); i++{
        num, _ := strconv.Atoi(string(s[i]))
        if num == 9{
            ans = append(ans, s[:i]+"0"+s[i+1:])
            ans = append(ans, s[:i]+"8"+s[i+1:])
            continue
        } else if num == 0{
            ans = append(ans, s[:i]+"1"+s[i+1:])
            ans = append(ans, s[:i]+"9"+s[i+1:])
            continue
        } else{
            ans = append(ans, s[:i]+strconv.Itoa(num+1)+s[i+1:])
            ans = append(ans, s[:i]+strconv.Itoa(num-1)+s[i+1:])
        }
    }

    return ans
}

func openLock(deadends []string, target string) int {
    m := make(map[string]struct{})
    for _, str := range deadends{
        m[str] = struct{}{}
    }

    start := "0000"
    
    if _, ok := m[start]; ok{
        return -1
    }

    queue := []string{start}
    steps := 0

    visited := make(map[string]struct{})
    visited[start] = struct{}{}
    for len(queue) > 0{
        size := len(queue)
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            if front == target{return steps}
            for _, neighbour := range getNeighbours(front){
                if _, ok := m[neighbour]; ok{continue}
                if _, ok := visited[neighbour]; ok{continue}
                queue = append(queue, neighbour)
                visited[neighbour] = struct{}{}
            }
        }
        steps++
    }

    return -1
}