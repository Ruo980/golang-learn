package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func Exec344() {
	// 填充式输出题目名称
	title := "反转字符串"
	utils.PrintName(title)

	//===开始输入相关参数
	// 执行输入操作
	var str string
	fmt.Println("请输入字符串s：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Scanln(&str)
	// 清空输入缓冲区
	reader.Discard(reader.Buffered())
	fmt.Println("开始‘反转字符串’函数")
	s := utils.GetArrayByString(str)
	reverseString(s)
	fmt.Println("结果为：")
	// 将字节byte转为字符输出
	utils.PrintName("‘反转字符串’程序结束")
}
