/**
  @author: dingfeng
  @date: 2021/1/12
  @note:
**/
package Week_02

import "sort"

//paixu
func repeategroupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _,str:=range strs{
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i]<s[j]
		})
		sortStr := string(s)
		hash[sortStr] = append(hash[sortStr],str)
	}
	var res [][]string
	for _,item:=range hash{
		res  = append(res,item)
	}
	return res
}

//计数
func repeategroupAnagrams2(strs []string) [][]string {
	hash := map[[26]int][]string{}
	for _,str:=range strs{
		 cnt := [26]int{}
		 for _,s:=range str{
		 	cnt[s-'a']++
		 }
		 hash[cnt] = append(hash[cnt],str)
	}
	var res [][]string
	for _,item:=range hash{
		res  = append(res,item)
	}
	return res
}