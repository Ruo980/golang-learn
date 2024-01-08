package function

func largestPerimeter(nums []int) int {
	maxLength := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j] > nums[k] && nums[i]+nums[k] > nums[j] && nums[j]+nums[k] > nums[i] && nums[i]-nums[j] < nums[k] && nums[i]-nums[k] < nums[j] && nums[j]-nums[k] < nums[i] {
					maxLength = max(maxLength, nums[i]+nums[j]+nums[k])
				}
			}
		}
	}
	return maxLength
}

func Exec976() {

}
