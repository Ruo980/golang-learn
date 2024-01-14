package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head != nil {
		L := head
		p := head.Next
		for p != nil {
			if p.Val != L.Val {
				L.Next = p
				L = p
			}
			p = p.Next
		}
		L.Next = nil
	}
	return head
}

func Exec83() {
	// 填充式输出题目名称
	title := "删除排序链表中的重复元素"
	utils.PrintName(title)

	// 接收参数
	var str string
	fmt.Println("请输入链表：")
	fmt.Scanln(&str)
	//转换为链表
	head := utils.GetList(str)
	// 执行算法
	result := deleteDuplicates(head)
	// 打印结果
	utils.PrintList(result)
	//算法结束
	utils.PrintName("‘删除排序链表中的重复元素’程序结束")
}
