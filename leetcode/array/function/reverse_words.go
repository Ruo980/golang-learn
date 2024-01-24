package function

import (
	"bufio"
	"fmt"
	"golang-learn/leetcode/array/function/utils"
	"os"
	"strings"
)

func reverseWords(s string) string {
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
	// 反转单词，去除中间多余空格
	i = 0  // 单词终点
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
			// j指向下一个单词的起点
			for j < len(str) && str[j] == ' ' {
				j++
			}
		}
		i++
	}
	return string(str[:k-1])
}

func Exec151() {
	// 填充式输出题目名称
	title := "翻转字符串里的单词"
	utils.PrintName(title)

	// 接收参数
	var str string
	fmt.Println("请输入字符串:")

	// 读取一行输入
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入失败:", err)
		return
	}
	// Windows下输入的字符串是\r\n结尾
	// 去除换行符
	str = strings.TrimRight(input, "\n")
	// 去除回车符
	str = strings.TrimRight(str, "\r")
	// 执行算法
	result := reverseWords(str)
	// 打印结果
	fmt.Println("翻转后的字符串为:", result)
	//算法结束
	utils.PrintName("‘翻转字符串里的单词’程序结束")
}
