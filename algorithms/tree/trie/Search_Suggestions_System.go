/*
You are given an array of strings products and a string searchWord.

Design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.

Return a list of lists of the suggested products after each character of searchWord is typed.
*/

func suggestedProducts(products []string, searchWord string) [][]string {
    slices.Sort(products)
    var ans [][]string

    for i := 0; i < len(searchWord); i++ {
        pattern := searchWord[:i+1]
        var group []string
        for _, p := range products {
            if strings.HasPrefix(p, pattern) {
                group = append(group, p)
                if len(group) == 3 {
                    break
                }
            }
        }
        ans = append(ans, group)
    }

    return ans
}