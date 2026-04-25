/*
A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.

Suppose we need to investigate a mutation from a gene string startGene to a gene string endGene where one mutation is defined as one single character changed in the gene string.

For example, "AACCGGTT" --> "AACCGGTA" is one mutation.
There is also a gene bank bank that records all the valid gene mutations. A gene must be in bank to make it a valid gene string.

Given the two gene strings startGene and endGene and the gene bank bank, return the minimum number of mutations needed to mutate from startGene to endGene. If there is no such a mutation, return -1.

Note that the starting point is assumed to be valid, so it might not be included in the bank.
*/

func getNeighbours(s string, m, visited map[string]struct{}) []string{
    var ans []string

    for i := 0; i < len(s); i++{
        for _, symbol := range []string{"A","C","G","T"}{
            if symbol != string(s[i]){
                ns := s[:i]+symbol+s[i+1:]
                if _, ok := m[ns]; ok{
                    if _, ok := visited[ns]; !ok{
                        visited[ns] = struct{}{}
                        ans = append(ans, ns)
                    }
                }
            }
        }
    }

    return ans
}

func minMutation(startGene string, endGene string, bank []string) int {
    m := make(map[string]struct{})
    for _, str := range bank{
        m[str] = struct{}{}
    }

    if _, ok := m[endGene]; !ok{
        return -1
    }

    queue := []string{startGene}
    steps := 0

    visited := make(map[string]struct{})
    visited[startGene] = struct{}{}
    for len(queue) > 0{
        size := len(queue)
        for i := 0; i < size; i++{
            front := queue[0]
            queue = queue[1:]
            if front == endGene{return steps}
            for _, neighbour := range getNeighbours(front, m, visited){
                queue = append(queue, neighbour)
            }
        }
        steps++
    }

    return -1
}