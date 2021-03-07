/**
  @author: dingfeng
  @date: 2021/3/5
  @note:
**/
package leetcode

func lengthOfLastWord(s string) int {
	maxindex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			maxindex = i
		}
	}
	return len(s)-maxindex-1
}
