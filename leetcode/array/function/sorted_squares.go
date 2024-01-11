package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func sortedSquares(nums []int) []int {
	length := len(nums)
	i := 0
	j := length - 1
	t := length - 1
	res := make([]int, length)
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[t] = nums[i] * nums[i]
			i++
		} else {
			res[t] = nums[j] * nums[j]
			j--
		}
		t--
	}
	return res
}

func Exec977() {
	// 填充式输出题目名称
	title := "有序数组的平方"
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
	fmt.Println("开始‘有序数组的平方’函数")
	sortedSquares(nums)
	fmt.Println("有序数组的平方为：", nums)
	utils.PrintName("‘有序数组的平方’程序结束")
}
