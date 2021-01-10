/**
  @author: dingfeng
  @date: 2021/1/10
  @note:
**/
package Week_02

import "sort"

// hash
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := map[string]int{}
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if _, ok := hash[ch]; ok {
			hash[ch]++
		} else {
			hash[ch] = 1
		}
	}
	for j := 0; j < len(t); j++ {
		ch := string(t[j])
		if _, ok := hash[ch]; ok {
			hash[ch]--
			if hash[ch] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// 方法二 排序
func isAnagram2(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}


func isAnagram3(s, t string) bool {
	var alphabet [26]int
	for _,item:=range s{
		alphabet[item-'a']++
	}
	for _,item:=range  t{
		alphabet[item-'a']--
	}
	for _,item:=range alphabet{
		if item!=0{
			return false
		}
	}
	return true
}

