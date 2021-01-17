/**
  @author: dingfeng
  @date: 2021/1/17
  @note:
**/
package Week_03

import "sort"

func PermuteUnique(nums []int) [][]int {
	if len(nums) == 0{
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	var path []int

	var dfs func(depth int, used []bool, path []int)
	dfs = func(depth int, used []bool, path []int) {
		if len(nums) == depth {
			//两种写法都可以
			//res = append(res, append([]int(nil), path...))
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {

			if used[i] || (i > 0 && nums[i] == nums[i-1] && used[i-1] == false) {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(depth+1, used, path)
			used[i] = false
			path = path[:len(path)-1]

		}
	}
	dfs(0, make([]bool, len(nums)), path)
	return res

}
