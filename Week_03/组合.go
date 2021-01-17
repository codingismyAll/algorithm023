/**
  @author: dingfeng
  @date: 2021/1/17
  @note:
**/
package Week_03

func Combine(n int, k int) [][]int {

	if n < 0 || k > n {
		return nil
	}

	var res [][]int
	var dfs func(n, k, begin int, path []int)
	dfs = func(n, k, begin int, path []int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := begin; i <= n; i++ {
			path = append(path, i)
			dfs(n, k, i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(n,k,1,[]int{})
	return res
}
