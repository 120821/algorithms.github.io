package main

import "fmt"

// 链表节点结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// 创建链表
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	curr := head

	for i := 1; i < len(vals); i++ {
		node := &ListNode{Val: vals[i]}
		curr.Next = node
		curr = node
	}

	return head
}

// 打印链表
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	// 创建链表
	vals := []int{1, 2, 3, 4, 5}
	head := createList(vals)

	fmt.Println("原始链表：")
	printList(head)

	// 反转链表
	reversed := reverseList(head)

	fmt.Println("反转后的链表：")
	printList(reversed)
}
