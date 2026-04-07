/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
Return the indices of the two numbers index1 and index2, each incremented by one, as an integer array [index1, index2] of length 2.
The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.
*/

func twoSum(numbers []int, target int) []int {
    l,r := 0,len(numbers)-1
    for l < r{
        tmp := numbers[l]+numbers[r]
        if tmp == target{
            return []int{l+1, r+1}
        } else if tmp > target{
            r--
        } else{
            l++
        }
    }
    
    return nil
}