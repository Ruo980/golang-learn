package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func crossingMiddle(nums []int) int {
	mid := len(nums) / 2
	// 从中间向左右两侧扩展，找到最大子数组
	max1 := nums[mid]
	sum1 := nums[mid]
	for i := mid - 1; i >= 0; i-- {
		sum1 += nums[i]
		if sum1 > max1 {
			max1 = sum1
		}
	}
	max2 := nums[mid]
	sum2 := nums[mid]
	for i := mid + 1; i < len(nums); i++ {
		sum2 += nums[i]
		if sum2 > max2 {
			max2 = sum2
		}
	}
	return max1 + max2 - nums[mid]
}

// 递归算法实现最大子数组和求解：最大子数组可能在左侧、右侧以及中间
func maxSubArray(nums []int) int {
	// 递归终止条件
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 0 {
		return -1
	} else {
		leftmax := maxSubArray(nums[:len(nums)/2])
		rightmax := maxSubArray(nums[len(nums)/2:])
		midmax := crossingMiddle(nums)
		return max(leftmax, rightmax, midmax)
	}
}

func Exec53() {
	// 填充式输出题目名称
	title := "Maximum Subarray"
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
	fmt.Println("开始‘最大子数组和’函数")
	max := maxSubArray(nums)
	fmt.Println("最大子数组和为：")
	fmt.Println(max)
	utils.PrintName("‘最大子数组和’程序结束")
}
