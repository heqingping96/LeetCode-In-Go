package leetcode_problems

/*将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil&&l2 != nil{
		return l2
	}
	if l2 == nil &&l1!= nil{
		return l1
	}
	if l1 == nil&&l2 == nil{
		return nil
	}
	var newhead *ListNode
	cur1 ,cur2:= l1,l2
	if cur1.Val<cur2.Val{
		newhead = cur1
		newhead.Next = mergeTwoLists(cur1.Next,l2)
		return l1
	}else {
		newhead = cur2
		newhead.Next = mergeTwoLists(cur1,cur2.Next)
		return l2
	}
}

