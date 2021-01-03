/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hash:= map[int]int{}
	for i,num:=range nums{
		if v,ok:=hash[target-num];ok{
			return []int{v,i}
		}
		hash[num] = i
	}
	return nil

}
// @lc code=end

