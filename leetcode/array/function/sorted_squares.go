package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
	"sort"
)

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
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
