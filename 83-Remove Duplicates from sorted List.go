package leetcode_problems

/*给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	dummnyhead := &ListNode{
		Val:-1,
		Next:head,
	}
	cur := dummnyhead
	for cur.Next!=nil&&cur.Next.Next != nil{
		if cur.Next.Val == cur.Next.Next.Val{
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummnyhead.Next
}


