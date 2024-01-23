package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
)

func reverseWords(s string) string {
	fmt.Println(s)
	str := []rune(s)
	i := len(str) - 1
	j := 0
	// 反转整个字符串
	for j < i {
		str[j], str[i] = str[i], str[j]
		j++
		i--
	}
	// 去除前后所有空格，剪枝操作
	i = 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	j = len(str) - 1
	for j >= 0 && str[j] == ' ' {
		j--
	}
	str = str[i : j+1]
	// 遍历
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	// 反转单词，去除中间多余空格
	i = 0
	j = 0  // 单词起点
	k := 0 // 新的数组
	// 增加一个尾部空格：使得最后一个单词反转和其他单词一样
	str = append(str, ' ')
	// 反转单词
	for i < len(str) {
		if str[i] == ' ' && str[i-1] != ' ' {
			// 一个单词扫描结束，将单词以及空格进行接收
			t := i - 1
			for j < t {
				str[j], str[t] = str[t], str[j]
				str[k] = str[j]
				j++
				t--
				k++
			}
			for j <= i {
				str[k] = str[j]
				k++
				j++
			}
			// 单词起点
			j = i + 1
		}
		i++
	}
	// 去除尾部空格
	return string(str[:k])
}

func Exec151() {
	// 填充式输出题目名称
	title := "翻转字符串里的单词"
	utils.PrintName(title)

	// 接收参数
	var str string
	fmt.Println("请输入字符串：")
	// 读取一行输入并丢弃
	reader := bufio.NewReader(os.Stdin)
	// 清空输入缓冲区
	_, err := reader.Discard(reader.Buffered())
	if err != nil {
		return
	}
	_, err1 := fmt.Scanln(&str)
	if err1 != nil {
		return
	}
	// 清空输入缓冲区
	_, err2 := reader.Discard(reader.Buffered())
	if err2 != nil {
		return
	}

	// 执行算法
	result := reverseWords(str)
	// 打印结果
	fmt.Println("翻转后的字符串为：", result)
	//算法结束
	utils.PrintName("‘翻转字符串里的单词’程序结束")
}
