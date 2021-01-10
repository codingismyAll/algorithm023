/**
  @author: dingfeng
  @date: 2021/1/10
  @note:
**/
package Week_02

// 暴力求解
func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			if nums[j]+nums[i] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}
// hash
func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if v,ok:=hash[target-nums[i]];ok{
			return []int{i,v}
		}
		hash[nums[i]] = i
	}
	return nil
}