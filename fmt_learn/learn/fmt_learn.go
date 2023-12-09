// @Program: golang-learn
// @Author: FunCoder
// @Create: 2023-09-03 16:30
// @Description: 字符串填充：在字符串前添加或附加空格或特殊字符的操作，这样无论输入字符串的长度如何，最终字符串的总长度都是固定的。
package learn

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func FmtLearn() {
	text := "Hello"
	padding := 18 - utf8.RuneCountInString(text) // 中文字符串一个占用3位，因此我们需要计算真实位数

	if padding > 0 {
		leftPadding := padding / 2
		rightPadding := padding - leftPadding
		paddedText := fmt.Sprintf("%s%s%s", strings.Repeat("=", leftPadding), text, strings.Repeat("=", rightPadding))
		fmt.Printf("%18s\n", paddedText)
	} else {
		fmt.Println(text)
	}
}
