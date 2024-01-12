// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-03 16:51
// @Description: 填充式输出字符串
package utils

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 给定一个字符串形式的名，按照填充的方式输出该字符串

func PrintName(title string) {
	padding := 40 - utf8.RuneCountInString(title) // 中文字符串一个占用3位，因此我们需要计算真实位数
	if padding > 0 {
		leftPadding := padding / 2
		rightPadding := padding - leftPadding
		paddedText := fmt.Sprintf("%s%s%s", strings.Repeat("=", leftPadding), title, strings.Repeat("=", rightPadding))
		fmt.Printf("%18s\n", paddedText)
	} else {
		fmt.Println(title)
	}
}
