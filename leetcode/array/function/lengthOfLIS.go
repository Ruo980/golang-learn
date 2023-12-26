package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

// 从nums数组中查询最大递增序列长度：比较各起点的最大递增序列长度，返回以i为起点的最大递增子序列长度
func max_length(a []int, nums []int, i int) int {
	if a[i] != 0 {
		return a[i]
	} else {
		if i == len(nums)-i {
			return 1
		} else {
			max_len := 1 // 最小长度为1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i] {
					max_len = max(max_len, 1+max_length(a, nums, j))
				}
			}
			a[i] = max_len
			return max_len
		}
	}
}

func lengthOfLIS(nums []int) int {
	a := make([]int, len(nums)) // 长度为len(nums)的数组，用于存储以i为起点的最大递增子序列长度，默认值为0
	max_nums := 1
	for i := 0; i < len(nums); i++ {
		max_nums = max(max_nums, max_length(a, nums, i))
	}
	return max_nums
}

func Exec300() {
	// 填充式输出题目名称
	title := "lengthOfLIS"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作：输入数组
	fmt.Println("请输入数组值：")
	var input string
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&input)
	// 输入格式为：[-2,1,-3,4,-1,2,1,-5,4]
	nums := utils.GetArray(input)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘最长递增子序列’函数")
	max_len := findLengthOfLCIS(nums)
	fmt.Println("最长子数组长度为：", max_len)
	utils.PrintName("‘最长递增子序列’程序结束")
}
