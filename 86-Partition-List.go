package leetcode_problems

/*给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil{
		return head
	}
	dummnyhead1 := &ListNode{
		Val:-1,
		Next:nil,
	}
	dummnyhead2 := &ListNode{
		Val: -1,
		Next: nil,
	}
	cur := head
	cur1 := dummnyhead1
	cur2 := dummnyhead2
	for cur != nil {
		if cur.Val<x{
			cur1.Next = cur
			cur = cur.Next
			cur1= cur1.Next
			cur1.Next = nil
		}else {
			cur2.Next = cur
			cur = cur.Next
			cur2= cur2.Next
			cur2.Next = nil
		}
	}
	if dummnyhead2.Next == nil{
		return dummnyhead1.Next
	}else {
		cur1.Next = dummnyhead2.Next
		return dummnyhead1.Next
	}
}
