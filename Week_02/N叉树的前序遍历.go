/**
  @author: dingfeng
  @date: 2021/1/10
  @note:
**/
package Week_02

type Node struct {
	Val      int
	Children []*Node
}

//dfs 迭代
func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			if len(root.Children) == 0 {
				break
			}
			for i:=len(root.Children)-1;i>0;i--{
				stack = append(stack,root.Children[i])
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

// 迭代
var res []int
func preorder2(root *Node) []int {
	res = []int{}
	dfs(root)
	return res
}
func dfs(root *Node)  {
	if root != nil {
		res = append(res, root.Val)
		for _,n := range root.Children {
			dfs(n)
		}
	}
}

