package leetcode

/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 */

// @lc code=start
func numIdenticalPairs(nums []int) int {
    var c int = 0
    
    for i:= range nums {
        for j := range nums {
            if nums[i] == nums[j] && i < j {
                c++
            }
        }
    }
    return c
}
// @lc code=end

