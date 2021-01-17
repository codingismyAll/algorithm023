/**
  @author: dingfeng
  @date: 2021/1/15
  @note:
**/
package main

import (
	"algorithm023/Week_03"
	"fmt"
)

func main(){
	//tn := Week_03.TreeNode{
	//	Val:   1,
	//	Left:  &Week_03.TreeNode{
	//		Val:   2,
	//		Left:  &Week_03.TreeNode{
	//			Val: 4,
	//			Left:  &Week_03.TreeNode{
	//				Val: 8,
	//			},
	//			Right: &Week_03.TreeNode{
	//				Val: 9,
	//			},
	//		},
	//		Right: &Week_03.TreeNode{
	//			Val: 5,
	//			Left:  &Week_03.TreeNode{
	//				Val: 10,
	//			},
	//			Right: &Week_03.TreeNode{
	//				Val: 11,
	//			},
	//		},
	//	},
	//	Right: &Week_03.TreeNode{
	//		Val:   3,
	//		Left:  &Week_03.TreeNode{
	//			Val: 6,
	//			Left:  &Week_03.TreeNode{
	//				Val: 12,
	//			},
	//			Right: &Week_03.TreeNode{
	//				Val: 13,
	//			},
	//		},
	//		Right: &Week_03.TreeNode{
	//			Val: 7,
	//			Left:  &Week_03.TreeNode{
	//				Val: 14,
	//			},
	//			Right: &Week_03.TreeNode{
	//				Val: 15,
	//			},
	//		},
	//	},
	//}
	//res:=Week_03.LowestCommonAncestor(&tn,&Week_03.TreeNode{Val: 9},&Week_03.TreeNode{Val: 11})
	res := Week_03.Combine(4,2)
	fmt.Println(res)
}