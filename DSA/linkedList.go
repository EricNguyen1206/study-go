package DSA

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

// 707. Design Linked List
type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, tail: nil, size: 0}
}

func (list *MyLinkedList) Get(index int) int {
	if index < 0 || index >= list.size {
		return -1
	}
	cur := list.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (list *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val, Next: list.head, Prev: nil}
	if list.head != nil {
		list.head.Prev = newNode
	}
	list.head = newNode
	if list.tail == nil {
		list.tail = newNode
	}
	list.size++
}

func (list *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val, Next: nil, Prev: list.tail}
	if list.tail != nil {
		list.tail.Next = newNode
	} else {
		list.head = newNode
	}
	list.tail = newNode
	list.size++
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > list.size {
		return
	}
	if index == 0 {
		list.AddAtHead(val)
	} else if index == list.size {
		list.AddAtTail(val)
	} else {
		prev := list.head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		newNode := &ListNode{Val: val, Next: prev.Next, Prev: prev}
		if prev.Next != nil {
			prev.Next.Prev = newNode
		}
		prev.Next = newNode
		list.size++
	}
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.size {
		return
	}
	if index == 0 {
		list.head = list.head.Next
		if list.head != nil {
			list.head.Prev = nil
		} else {
			list.tail = nil
		}
	} else {
		prev := list.head
		for i := 0; i < index-1; i++ {
			prev = prev.Next
		}
		if prev.Next != nil {
			prev.Next = prev.Next.Next
			if prev.Next != nil {
				prev.Next.Prev = prev
			} else {
				list.tail = prev
			}
		}
	}
	list.size--
}

// 876. Middle of the Linked List
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 83. Remove Duplicates from Sorted List
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 203. Remove Linked List Elements
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 160. Intersection of Two Linked Lists
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

// 21. Merge Two Sorted Lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// 234. Palindrome Linked List
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow, slow.Next, prev = slow.Next, prev, slow
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow, prev = slow.Next, prev.Next
	}
	return true
}

// 146. LRU Cache
type LRUCache struct {
	cache    map[int]*ListNode
	head     *ListNode
	tail     *ListNode
	capacity int
}

func LRUConstructor(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*ListNode), head: &ListNode{}, tail: &ListNode{}, capacity: capacity}
}

func (list *LRUCache) Get(key int) int {
	if node, ok := list.cache[key]; ok {
		list.moveToHead(node)
		return node.Val
	}
	return -1
}

func (list *LRUCache) Put(key int, value int) {
	if node, ok := list.cache[key]; ok {
		node.Val = value
		list.moveToHead(node)
		return
	}
	node := &ListNode{Val: value}
	list.cache[key] = node
	list.moveToHead(node)
	if list.capacity < len(list.cache) {
		list.removeTail()
	}
}

func (list *LRUCache) moveToHead(node *ListNode) {
	if node == list.head.Next {
		return
	}
	list.removeNode(node)
	list.addNodeToHead(node)
}

func (list *LRUCache) removeNode(node *ListNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (list *LRUCache) addNodeToHead(node *ListNode) {
	node.Prev = list.head
	node.Next = list.head.Next
	list.head.Next.Prev = node
	list.head.Next = node
}

func (list *LRUCache) removeTail() {
	node := list.tail.Prev
	list.removeNode(node)
	delete(list.cache, node.Val)
}
