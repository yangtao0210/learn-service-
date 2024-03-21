package interview

// FindMaxConsecutiveOnes 描述：给定一个二进制数组 nums ， 计算其中最大连续 1 的个数
func FindMaxConsecutiveOnes(nums []int) int {
	res := 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 1 {
			res = max(res, right-left)
			left = right + 1
		}
		right++
	}
	res = max(res, right-left)
	return res
}
