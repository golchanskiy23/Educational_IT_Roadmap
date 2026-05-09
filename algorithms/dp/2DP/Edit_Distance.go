/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character
*/

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1)+1)
    for i := 0; i <= len(word1); i++{
        dp[i] = make([]int, len(word2)+1)
    }

    for j := 1; j <= len(word2); j++{
        dp[0][j] = j
    }

    for i := 1; i <= len(word1); i++{
        dp[i][0] = i
    }

    for i := 1; i <= len(word1); i++{
        for j := 1; j <= len(word2); j++{
            if word1[i-1] == word2[j-1]{
                dp[i][j] = dp[i-1][j-1]
                continue
            }

            dp[i][j] = min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1]))+1
        }
    } 
    // insert: dp[i+1][j]
    // delete: dp[i-1][j]
    // replace: dp[i-1][j-1]
    // s1[i] = s2[j] continue
    return dp[len(word1)][len(word2)]
}