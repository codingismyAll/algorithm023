/**
  @author: dingfeng
  @date: 2021/1/28
  @note:
**/
package leetcode

//dfs
func generateParenthesis(n int) []string {
	var res []string
	var dfs func(left, right, n int, s string)
	dfs = func(left, right, n int, s string) {
		if left == n && right == n {
			res = append(res, s)
			return
		}
		if left < n {
			dfs(left+1, right, n, s+"(")
		}
		if right < left {
			dfs(left, right+1, n, s+")")
		}
	}
	dfs(0, 0, n, "")
	return res
}

type Node struct {
	left  int
	right int
	str   string
}

//bfs
func GenerateParenthesis2(n int) []string {
	var res []string
	queue := []*Node{{
		left:  0,
		right: 0,
		str:   "",
	}}
	for len(queue) > 0 {
		lenth := len(queue)
		for i := 0; i < lenth; i++ {
			currnode := queue[i]
			if currnode.left == n && currnode.right == n {
				res = append(res, currnode.str)
			}
			if currnode.left < n {
				queue = append(queue, &Node{
					left:  currnode.left + 1,
					right: currnode.right,
					str:   currnode.str + "(",
				})
			}
			if currnode.right < currnode.left {
			//if currnode.right < currnode.left && currnode.right < n {
				queue = append(queue, &Node{
					left:  currnode.left,
					right: currnode.right + 1,
					str:   currnode.str + ")",
				})
			}

		}
		queue = queue[lenth:]
	}
	return res
}
