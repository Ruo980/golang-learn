// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-02 20:41
// @Description: 两数之和
// @input [1,] [3,4]
package function

import (
	"bufio"
	"fmt"
	"os"

	"golang-learn/leetcode/array/function/utils"
)

// twoSum
//
//	@Description: 两数之和
//	@param nums
//	@param target
//	@return []int
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if m[target-nums[i]] != 0 {
			return []int{m[target-nums[i]] - 1, i}
		}
		m[nums[i]] = i + 1
	}
	return nil
}

func Exec1() {
	// 填充式输出题目名称
	title := "两数之和"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var nums []int
	var str string
	var target int
	fmt.Println("请输入数组nums：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	_, err3 := reader.Discard(reader.Buffered())
	if err3 != nil {
		return
	}
	_, err2 := fmt.Scanln(&str)
	if err2 != nil {
		return
	}
	nums = utils.GetArray(str)
	fmt.Println("请输入target：")
	_, err := fmt.Scanln(&target)
	if err != nil {
		return
	}
	// 清空输入缓冲区
	_, err3 = reader.Discard(reader.Buffered())
	if err3 != nil {
		return
	}
	fmt.Println("开始‘两数之和’函数")
	result := twoSum(nums, target)
	fmt.Println("结果为：", result)
	utils.PrintName("‘两数之和’程序结束")
}
