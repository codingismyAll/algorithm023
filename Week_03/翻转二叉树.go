/**
  @author: dingfeng
  @date: 2021/1/15
  @note:
**/
package Week_03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


//dfs
func InvertTree(root *TreeNode) *TreeNode {
	//terminator
	if root == nil{
		return nil
	}
	//process logic
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	//drill down
	InvertTree(root.Left)
	InvertTree(root.Right)
	return  root
	//reverse state
}

//bfs
func InvertTree2(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) == 0 {
		tmp:=queue[0]
		queue = queue[1:]
		left := tmp.Left
		tmp.Left = tmp.Right
		tmp.Right = left
		if tmp.Left != nil{
			queue = append(queue, tmp.Left)
		}
		if tmp.Right != nil{
			queue = append(queue,tmp.Right)
		}
	}
	return  root


}


