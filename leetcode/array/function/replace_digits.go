package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func replaceDigits(s string) string {
	resultRune := []rune(s)
	for i := 1; i < len(s); i += 2 {
		resultRune[i] = resultRune[i-1] + rune(resultRune[i]-'0')
	}
	return string(resultRune)
}

func Exec1844() {
	// 填充式输出题目名称
	title := "将所有数字用字符替换"
	utils.PrintName(title)

	// 接收参数
	var str string
	fmt.Println("请输入字符串：")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	// 执行算法
	result := replaceDigits(str)
	// 打印结果
	fmt.Println("替换后的字符串为：", result)
	//算法结束
	utils.PrintName("‘将所有数字用字符替换’程序结束")
}
