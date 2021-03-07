/**
  @author: dingfeng
  @date: 2021/3/5
  @note:
**/
package leetcode

import "strings"

func toLowerCase(str string) string {
	return  strings.ToLower(str)
}
func toLowerCase2(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for i := 0; i < len(str); i++ {
		c := str[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b.WriteByte(c)
	}
	return b.String()
}