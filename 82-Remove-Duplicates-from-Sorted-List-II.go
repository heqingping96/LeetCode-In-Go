package leetcode_problems

//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
/*
示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
 */

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil{
		return head
	}

	dummnyhead := &ListNode{
		Val:-1,
		Next:head,
	}

	cur :=dummnyhead
	for cur.Next != nil &&cur.Next.Next!= nil{
		if cur.Next.Val == cur.Next.Next.Val{
			tmp := cur.Next.Val
			for cur.Next != nil&&cur.Next.Val == tmp{
				cur.Next = cur.Next.Next
			}
		}else {
			cur = cur.Next
		}
	}
	return dummnyhead.Next
}


