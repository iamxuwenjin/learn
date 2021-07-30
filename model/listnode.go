package model

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func Constructor() MyLinkedList {
	var myLinkedList MyLinkedList
	return myLinkedList
}

func (listNode *MyLinkedList) Get(idx int) int {
	if idx >= listNode.size || idx < 0 {
		return -1
	}
	cur := 0
	node := listNode.Head
	var ans int
	for ; cur <= idx; cur++ {
		ans = node.Val
		node = node.Next
	}
	return ans
}

func (listNode *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: listNode.Head,
	}
	listNode.Head = node
	if listNode.Tail == nil {
		listNode.Tail = node
	}
	listNode.size++
}

func (listNode *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
	}
	if listNode.Tail != nil {
		listNode.Tail.Next = node
	}
	listNode.Tail = node
	if listNode.Head == nil {
		listNode.Head = node
	}
	listNode.size++
}

func (listNode *MyLinkedList) AddAtIndex(index int, val int) {
	cur := 0
	curNode := listNode.Head
	preNode := listNode.Head
	for ; cur <= index; cur++ {
		if cur == index {
			if index == 0 {
				listNode.AddAtHead(val)
			} else if index == listNode.size {
				listNode.AddAtTail(val)
			} else {
				node := &Node{
					Val:  val,
					Next: curNode,
				}
				preNode.Next = node
			}
			listNode.size++
			continue
		}
		preNode = curNode
		curNode = curNode.Next
	}
}

func (listNode *MyLinkedList) DeleteAtIndex(index int) {
	cur := 0
	curNode := listNode.Head
	preNode := listNode.Head
	for ; cur <= index; cur++ {
		if cur == index {
			if index == 0 {
				listNode.Head = curNode.Next
			} else if index == listNode.size-1 {
				preNode.Next = nil
				listNode.Tail = preNode
			} else {
				preNode.Next = curNode.Next
			}
			listNode.size--
			continue
		}
		preNode = curNode
		curNode = curNode.Next
	}
}

func (listNode *MyLinkedList) append(num int) {
	node := &Node{
		Val:  num,
		Next: nil,
	}
	if listNode.Head == nil && listNode.Tail == nil {
		listNode.Head = node
		listNode.Tail = node
	}
	if listNode.Tail != nil {
		listNode.Tail.Next = node
		listNode.Tail = node
	}
}

func GetListNode(nums []int) *MyLinkedList {
	var listNode MyLinkedList
	for _, num := range nums {
		listNode.append(num)
	}
	return &listNode
}
