package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	} else {
		i := 1 //i为快指针
		j := 0 //j为慢指针
		for i < len(nums) {
			if nums[i] > nums[j] {
				j++ // 待插入的索引位置
				nums[j] = nums[i]
			}
			i++
		}
		return j + 1
	}
}

func Exec26() {
	// 填充式输出题目名称
	title := "删除排序数组中的重复项"
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
	fmt.Println("开始‘删除排序数组中的重复项’函数")
	len := removeDuplicates(nums)
	fmt.Println("删除重复项后的数组为：", nums[:len])
	utils.PrintName("‘删除排序数组中的重复项’程序结束")
}
