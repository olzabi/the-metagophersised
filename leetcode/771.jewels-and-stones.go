package leetcode

/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 */

// @lc code=start
func numJewelsInStones(jewels string, stones string) int {
    // jewel = type | stones = counter
    // how many counter of type
    var c int
    var m = make(map[byte]int)
        
    for s := range stones { 
        m[stones[s]]++
    }
    
    for j := range jewels {
        c+=m[jewels[j]]
    }
    return c
}
// @lc code=end

