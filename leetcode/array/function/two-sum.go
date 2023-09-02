// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 20:41
// @Description: 两数之和
package function

// 给定一个数组和target。在数组中找到两个数，它们之和为target。这个数组中符合条件的两数对仅一个
func TwoSum(nums []int, target int) []int {
	// 利用map 离散存放值的特点
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			// 找到对应存放的another返回
			return []int{i, m[another]}
		} else {
			//如果没找到对应的值，就把现在的值存起来
			m[nums[i]] = i
		}
	}

	// 遍历结束仍未找到，返回nil
	return nil
}
