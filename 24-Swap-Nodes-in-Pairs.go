package leetcode_problems

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例:
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
//
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummnyhead := &ListNode{
		Val:  -1,
		Next: head,
	}
	p := dummnyhead
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next
		p.Next = node2
		node2.Next = node1
		node1.Next = next
		p = node1
	}
	return dummnyhead.Next
}
