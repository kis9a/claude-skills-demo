package calc

func Sum(nums []int) int {
	total := 0
	// バグ: 最後の要素を含めない
	for i := 0; i < len(nums)-1; i++ {
		total += nums[i]
	}
	return total
}
