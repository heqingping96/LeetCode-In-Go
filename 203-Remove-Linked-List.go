package leetcode_problems

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

 */

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	dummnyhead := &ListNode{
		Val :-1,
		Next :head,
	}
	cur := dummnyhead
	for cur.Next != nil{
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummnyhead.Next
}