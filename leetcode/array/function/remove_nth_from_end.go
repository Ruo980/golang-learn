package function

import (
	"fmt"
	"golang-learn/leetcode/array/function/utils"
)

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	// 本题的关键就是建立一个虚拟头节点：因为可能删除的是头节点，所以需要一个虚拟头节点，即凡是涉及到头节点的操作，都需要用到虚拟头节点
	p1 := head
	p2 := &utils.ListNode{}
	p2.Next = head
	p3 := p2
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p3 = p3.Next
	}
	p3.Next = p3.Next.Next
	return p2.Next
}

func Exec19() {
	// 填充式输出题目名称
	title := "删除链表的倒数第N个节点"
	utils.PrintName(title)

	// 接收参数
	var str string
	var n int
	fmt.Println("请输入链表：")
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}
	fmt.Println("请输入n：")
	_, err = fmt.Scanln(&n)
	if err != nil {
		return
	}
	//转换为链表
	head := utils.GetList(str)
	// 执行算法
	result := removeNthFromEnd(head, n)
	// 打印结果
	utils.PrintList(result)
	//算法结束
	utils.PrintName("‘删除链表的倒数第N个节点’程序结束")
}
