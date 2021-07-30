package leetcode

import "learn/model"

func detectCycle(head *model.Node) *model.Node {
	ansMap := make(map[*model.Node]bool)
	for head != nil {
		if _, ok := ansMap[head]; ok {
			return head
		}
		ansMap[head] = true
		head = head.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var ans *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = ans
		ans = curNode
		curNode = nextNode
	}
	return ans
}
