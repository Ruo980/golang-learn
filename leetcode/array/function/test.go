package function

func test(nums []int) int {
	if len(nums) == 1 {
		return 1
	} else {
		i := 1 //i为快指针
		j := 0 //j为慢指针
		for i < len(nums) {
			if nums[i] > nums[j] {
				j++ // 待插入的索引位置
				nums[j] = nums[i]
			}
			i++
		}
		return j + 1
	}
}
