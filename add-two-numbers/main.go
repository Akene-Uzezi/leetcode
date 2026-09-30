package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1arr := make([]int, lenlist(l1))
	l1arr = inputValueToArray(l1arr, l1)
	return nil
}

func inputValueToArray(arr []int, l *ListNode) []int {
	current := l
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}

func lenlist(head *ListNode) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func main() {
	testl := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	fmt.Println(testl.Val, testl.Next.Val, testl.Next.Next.Val, testl.Next.Next.Next)
	fmt.Println(lenlist(&testl))
}
