package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1arr := make([]int, lenlist(l1))
	l2arr := make([]int, lenlist(l2))
	l1int := ivra(l1arr, l1)
	l2int := ivra(l2arr, l2)
	sum := l1int + l2int
	return itarl(sum)
}

func reverseArrayIntoInt(arr []int) int {
	count := 0
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	for _, digit := range arr {
		count = count*10 + digit
	}
	return count
}

func ivra(arr []int, l *ListNode) int {
	narr := inputValueToArray(arr, l)
	return reverseArrayIntoInt(narr)
}

func inputValueToArray(arr []int, l *ListNode) []int {
	current := l
	i := 0
	for current != nil {
		arr[i] = current.Val
		i++
		current = current.Next
	}
	return arr
}

func intToArr(num int) []int {
	if num == 0 {
		return []int{0}
	}

	var digits []int
	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num = num / 10
	}
	return digits
}

func itarl(integer int) *ListNode {
	intarr := intToArr(integer)
	for i, j := 0, len(intarr)-1; i < j; i, j = i+1, j-1 {
		intarr[i], intarr[j] = intarr[j], intarr[i]
	}
	return arrayToLinkedList(intarr)
}

func arrayToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
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
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	result := addTwoNumbers(l1, l2)
	fmt.Println(result, result.Val, result.Next, result.Next.Next, result.Next.Next.Next)
}
