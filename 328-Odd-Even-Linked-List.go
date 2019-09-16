package leetcode_problems

/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
 */

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	dummnyhead1:= &ListNode{
		Val:-1,
		Next:nil,
	}
	dummnyhead2:= &ListNode{
		Val:-1,
		Next:nil,
	}
	cur := head
	cur1,cur2 := dummnyhead1,dummnyhead2
	for i:=0;cur != nil;i++{
		if i%2 == 0{
			cur1.Next = cur
			cur = cur.Next
			cur1 = cur1.Next
			cur1.Next = nil

		}else {
			cur2.Next = cur
			cur = cur.Next
			cur2 = cur2.Next
			cur2.Next = nil

		}
	}
	if dummnyhead2.Next == nil&&dummnyhead1.Next!= nil{
		return dummnyhead1.Next
	}else if dummnyhead1.Next == nil&&dummnyhead2.Next!= nil{
		return dummnyhead2.Next}else {
		cur1.Next = dummnyhead2.Next
		return dummnyhead1.Next
	}
}
