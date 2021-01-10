/**
  @author: dingfeng
  @date: 2021/1/10
  @note:
**/
package Week_02

import "sort"

//排序
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr],str)
	}
	ans:=[][]string{}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

//计数
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _,str:=range strs{
		cnt := [26]int{}
		for _,b:=range str{
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt],str)
	}
	ans:=[][]string{}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans

}