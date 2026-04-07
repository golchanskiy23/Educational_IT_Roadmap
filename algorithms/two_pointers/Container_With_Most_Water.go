/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container
*/

func min(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func maxArea(height []int) int {
    l,r := 0, len(height)-1
    area := math.MinInt
    for l < r{
        area = max(area, (r-l)*min(height[l], height[r]))
        if height[l] < height[r]{
            l++
        } else{
            r--
        }
        fmt.Println(area)
    }
    return area
}