package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

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
	// 填充式输出题目名称
	title := "三角形的最大周长"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var nums []int
	var str string
	fmt.Println("请输入数组nums：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&str)
	nums = utils.GetArray(str)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘三角形的最大周长’函数")
	len := largestPerimeter(nums)
	fmt.Println("三角形的最大周长为：", len)
	utils.PrintName("‘三角形的最大周长’程序结束")
}
