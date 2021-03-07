package Week_08
func reverseBits(num uint32) uint32 {
	var bitsize uint = 31
	var ans uint32
	for num != 0{
		ans += (num & 1 << bitsize)
		bitsize--
		num = num >> 1
	}
	return ans
}
