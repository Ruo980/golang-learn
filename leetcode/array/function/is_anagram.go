package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func isAnagram(s string, t string) bool {
	var nums [26]int
	for _, v := range s {
		nums[v-'a']++
	}
	for _, v := range t {
		nums[v-'a']--
	}
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}

func Exec242() {
	// 填充式输出题目名称
	title := "有效的字母异位词"
	utils.PrintName(title)

	// 接收参数
	var s, t string
	fmt.Println("请输入s：")
	_, err := fmt.Scanln(&s)
	if err != nil {
		return
	}
	fmt.Println("请输入t：")
	_, err = fmt.Scanln(&t)
	if err != nil {
		return
	}
	// 执行算法
	result := isAnagram(s, t)
	// 打印结果
	fmt.Println(result)
	//算法结束
	utils.PrintName("‘有效的字母异位词’程序结束")
}
