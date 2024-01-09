package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func minSubArrayLen(target int, nums []int) int {
	i := 0      // i为慢指针
	j := 0      // j为快指针
	length := 0 // len为数组长度
	sum := 0    // sum为数组和
	minLen := -1
	for j < len(nums) {
		if sum < target {
			sum += nums[j]
			length++
			j++
			// 将nums[j]加入到数组中后，判断数组和是否大于target
			if sum >= target {
				if minLen == -1 {
					minLen = length
				}
				minLen = min(minLen, length)
			}
		} else {
			sum -= nums[i]
			length--
			i++
			// 将nums[i]移除数组后，判断数组和是否小于target
			if sum < target {
				minLen = min(minLen, length+1)
			}
		}
	}
	for i < len(nums) {
		sum -= nums[i]
		length--
		i++
		if sum < target {
			minLen = min(minLen, length+1)
			break
		}
	}
	if minLen == -1 {
		minLen = 0
	}
	return minLen
}

func Exec209() {
	// 填充式输出题目名称
	title := "长度最小的子数组"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var nums []int
	var target int
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
	fmt.Println("请输入target：")
	// 读取一行输入并丢弃
	reader2 := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Scanln(&target)
	// 清空输入缓冲区
	reader.Discard(reader2.Buffered())
	fmt.Println("开始‘长度最小的子数组’函数")
	len := minSubArrayLen(target, nums)
	fmt.Println("长度最小的子数组为：", len)
	utils.PrintName("‘长度最小的子数组’程序结束")
}
