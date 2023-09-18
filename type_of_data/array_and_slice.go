package main

import "fmt"

// main
//
//	@Description: 主要研究数组和切片，涉及初始化、赋值、添加等
func main() {

	fmt.Println("开始数组探究")
	// 数组探究
	var a [20]int     //数组定义时必须声明长度，会自动初始化过程
	fmt.Println(a[2]) //默认值为 0
	for i := 0; i < len(a); i++ {
		a[i] = i + 3
	}
	for index, value := range a {
		fmt.Println(index, value)
	}

	fmt.Println("开始切片探究")
	// 切片探究
	var b = make([]int, 20) //切片初始化使用 make 关键字，它会根据指定的数组类型和长度去创建一个长度，该长度可以动态变化。
	fmt.Println(b[2])       //默认值为 0
	for i := 0; i < len(b); i++ {
		b[i] = i + 3
	}
	for index, value := range b {
		fmt.Println(index, value)
	}
}
