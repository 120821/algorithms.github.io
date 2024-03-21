// 链表是通过结构和指针实现的。是一种线性集合，元素按照顺序存放，但是不像数组一样连续存储，每个元素节点包含2个基本单元： 数据和向下的一个节点的指针。
// 定义链表
package main

import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

// 创建链表
func createList(vals []int) *ListNode {
  head := &ListNode{}
  current := head
  for _, val := range vals {
    current.Next = &ListNode{Val: val}
    current = current.Next
  }
  return head.Next
}

// 打印链表
func printList(head *ListNode) {
  for head != nil {
    fmt.Println(head.Val, " ")
    head = head.Next
  }
  fmt.Println()
}

// 在尾部增加元素
func addAtTail(head *ListNode, val int) *ListNode{
  newNode := &ListNode{Val: val}
  if head == nil {
    return newNode
  }
  current := head
  for current.Next != nil {
    current = current.Next
  }
  current.Next = newNode
  return head
}

// 删除元素
func deleteNode(head *ListNode, val int) *ListNode{
  if head == nil {
    return nil
  }
  if head.Val == val {
    return head.Next
  }
  current := head
  for current.Next != nil {
    if current.Next.Val == val {
      current.Next = current.Next.Next
      return head
    }
    current = current.Next
  }
  return head
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
  var prev *ListNode = nil
  current := head
  for current != nil {
    nextTemp := current.Next
    current.Next = prev
    prev = current
    current = nextTemp
  }
  return prev
}

// 寻找中节点
func findMiddle(head *ListNode) *ListNode {
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }
  return slow
}
// 是否有环
func hasCycle(head *ListNode) bool {
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast {
      return true
    }
  }
  return false
}

// 排序链表
func sort(head *ListNode) *ListNode{
  if head == nil || head.Next == nil {
    return head
  }
  mid := findMiddle(head)
  left := head
  right := mid.Next
  mid.Next = nil
  return merge(sort(left), sort(right))
}
// 归并排序合并
func merge(left *ListNode , right *ListNode) *ListNode {
  dummy := &ListNode{}
  curr := dummy
  for left != nil && right == nil {
    if left.Val < right.Val {
      curr.Next = left
      left = left.Next
    } else {

      curr.Next = right
      right = right.Next
    }
    curr = curr.Next
  }

  return dummy.Next
}

func main() {
  arr := []int{5, 4, 2, 22, 44}
  listNode := createList(arr)
  printList(listNode)
  isHasCycle := hasCycle(listNode)
  mid := findMiddle(listNode)
  fmt.Println(isHasCycle)
  fmt.Println(mid)
  listNodeReserve := reverseList(listNode)
  fmt.Println(listNodeReserve)
  printList(listNodeReserve)
	sortedListNode := sort(listNode)
	printList(sortedListNode)
}

