package Week_08
func hammingWeight(num uint32) int {
	count:=0
	for num!=0{
		count++
		//  去除最低位的1
		num = num & (num-1)
	}
	return count
}
