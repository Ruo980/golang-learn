package function

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"golang-learn/leetcode/array/function/utils"
)

// threeSum
//
//	@Description: 三数之和：三层循环
//	@param nums
//	@return [][]int
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	// nums 排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 由于已经排好序了，所以不要循环而是双指针协作:有序数组下对撞指针
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			// 防止循环中出现 j == k
			if j == k {
				break
			} else {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

// Exec15
//
//	@Description: 执行 15.三数之和 题目程序
func Exec15() {
	// 填充式输出题目名称
	title := "3Sum"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	fmt.Println("请输入数组值：")
	var input string
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&input)
	nums := utils.GetArray(input)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘三数之和’函数")
	result := threeSum(nums)
	fmt.Println("结果为：")
	fmt.Println(result)
	utils.PrintName("‘三数之和’程序结束")
}
