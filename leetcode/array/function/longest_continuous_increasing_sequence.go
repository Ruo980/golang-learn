package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 如果不为1，则最大长度最少为1
	max := 1
	count := 1
	// 巧妙地避开了 i+1 是否越界的讨论
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			count++
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}

func Exec674() {
	// 填充式输出题目名称
	title := "Longest continuous increasing sequence"
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
	fmt.Println("开始‘最长连续递增序列’函数")
	max := findLengthOfLCIS(nums)
	fmt.Println("最长子数组长度为：", max)
	utils.PrintName("‘最长连续递增序列’程序结束")
}
