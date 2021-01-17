/**
  @author: dingfeng
  @date: 2021/1/17
  @note:
**/
package Week_03

//dfs
func Permute(nums []int) [][]int {
	//append 应该要好好去学学
	var res [][]int
	var path []int
	var dfs func(depth int,used []bool,path []int)
	dfs = func(depth int, used []bool, path []int) {
		if len(nums) == depth{
			var temp []int
			temp = append(temp,path...)
			res = append(res, temp)
			return
		}
		for i:=0;i<len(nums);i++ {
			if used[i]{
				continue
			}
			used[i] = true
			//path[i] = nums[i] 肯定不能这样写因为下面减去长度了
			path = append(path,nums[i])
			dfs(depth+1,used,path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs(0,make([]bool,len(nums)),path)
	return res
}