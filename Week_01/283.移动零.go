/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package week01
// @lc code=start
func moveZeroes(nums []int)  {
	a := 0
	for _,item:=range nums{
		if item != 0{
			nums[a] = item
			a++
		}
	}

	for i:=a;i<len(nums);i++{
		nums[a] = 0
	}
}
// @lc code=end


//解法2

func moveZeroes2(nums []int)  {
	j := 0
	for i:=0;i<len(nums);i++{
		if nums[i] != 0{
			nums[j] = nums[i]
			if i != j{
				nums[i] = 0
			}
			j++
		}
	}
}