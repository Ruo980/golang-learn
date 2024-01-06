package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func removeElement(nums []int, val int) int {
	i := 0 //i为快指针
	j := 0 //j为慢指针
	for i < len(nums) {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}

func Exec27() {
	// 填充式输出题目名称
	title := "移除元素"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var nums []int
	var val int
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
	fmt.Println("请输入要移除的元素val：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&val)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Println("开始‘移除元素’函数")
	len := removeElement(nums, val)
	fmt.Println("移除元素后的数组为：", nums[:len])
	utils.PrintName("‘移除元素’程序结束")
}
