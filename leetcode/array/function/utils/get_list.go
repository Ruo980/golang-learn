package utils

import "fmt"

// 链表结构体：我们假设链表头节点有值，链表结尾指向nil
type ListNode struct {
	Val  int
	Next *ListNode
}

// GetList
//
//	@Description: 输入参数格式为[1,2,3,4,5]字符串。输出为链表
//	@param str
//	@return *ListNode
func GetList(str string) *ListNode {
	nums := GetArray(str)
	fmt.Println("接到一个链表：", nums)
	// 创建一个头结点
	head := &ListNode{}
	// 创建一个指针指向头结点
	cur := head
	// 将数组中的元素逐个添加到链表中
	for _, v := range nums {
		// 创建一个新节点
		newNode := &ListNode{Val: v}
		// 将新节点添加到链表中
		cur.Next = newNode
		// 将指针指向新节点
		cur = cur.Next
	}
	// 返回头结点
	return head.Next
}

// PrintList
//
//	@Description: 打印链表
//	@param head
func PrintList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println("nil")
}
