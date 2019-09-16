package leetcode_problems

//给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
//
//
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
// 进阶:
//
// 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
// 示例:
//
//
//输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出: 7 -> 8 -> 0 -> 7
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
func Reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2 := Reverse(l1), Reverse(l2)
	ca := 0              //进位
	n, n1, n2 := 0, 0, 0 //n 表示两数相加的和，n1表示l1的值，n2表示l2的值
	var pre *ListNode
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			n1 = 0
		} else {
			n1 = cur1.Val
		}
		if cur2 == nil {
			n2 = 0
		} else {
			n2 = cur2.Val
		}
		n = n1 + n2 + ca
		node := &ListNode{
			Val:  n % 10,
			Next: pre,
		}
		pre = node
		ca = n / 10
		if cur1 != nil && cur1.Next != nil {
			cur1 = cur1.Next
		} else {
			cur1 = nil
		}
		if cur2 != nil && cur2.Next != nil {
			cur2 = cur2.Next
		} else {
			cur2 = nil
		}
	}

	if ca == 1 {
		node1 := &ListNode{
			Val:  1,
			Next: pre,
		}
		return node1
	}
	return pre
}
