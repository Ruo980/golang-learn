package function

import (
	"bufio"
	"fmt"
	"os"

	"golang-learn/leetcode/array/function/utils"
)

// maxAreaMe
//
//	@Description: 本人使用暴力解法，使用双循环的机制遍历所有可能寻找最大值并记录，这种做法会超时。
//	@param height
//	@return int
func maxAreaMe(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				// 实现逻辑：max = ((height[i]-height[j])*(j-i)>max)?((height[i]-height[j])*(j-i)):max
				if height[j]*(j-i) > max {
					max = height[j] * (j - i)
				}
			} else {
				if height[i]*(j-i) > max {
					max = height[i] * (j - i)
				}
			}

		}
	}
	return max
}

// maxArea
//
//	@Description: 标准答案：利用对撞指针的原理进行快速接题。
//	@param height
//	@return int
func maxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		width := j - i
		high := 0
		if height[i] > height[j] {
			high = height[j]
			j--
		} else {
			high = height[i]
			i++
		}
		if max < width*high {
			max = width * high
		}
	}
	return max
}

func Exec11() {
	// 填充式输出题目名称
	title := "Container With Most Water"
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
	fmt.Println("开始‘盛最多水容器’函数")
	max := maxArea(nums)
	fmt.Println("最大值为：")
	fmt.Println(max)
	utils.PrintName("‘盛最多水容器’程序结束")
}
