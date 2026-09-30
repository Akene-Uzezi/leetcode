package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1arr := make([]int, lenlist(l1))
	_ = ivra(l1arr, l1)
	return nil
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
	arr := []int{1, 2, 3, 4, 5}
	int := reverseArrayIntoInt(arr)
	fmt.Println(int)
}
