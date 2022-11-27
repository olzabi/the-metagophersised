package leetcode

/*
 * @lc app=leetcode id=1365 lang=golang
 *
 * [1365] How Many Numbers Are Smaller Than the Current Number
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
    var res []int
    for _, i := range nums {
        c := 0
        for _, j := range nums {
            if j < i {
                c++
            }
        } 
        res = append(res, c)
    }
    return res 
}
// @lc code=end

