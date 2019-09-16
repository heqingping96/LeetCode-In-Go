package leetcode_problems

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1, cur2 := l1, l2
	n, n1, n2 := 0, 0, 0 // n表示两者相加，n1表示l1的值，n2表示n2的值
	ca := 0
	dummnyhead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cu := dummnyhead

	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			n1 = 0
		} else {
			n1 = cur1.Val
		}
		if cur2 != nil {
			n2 = cur2.Val
		} else {
			n2 = 0
		}
		n = n1 + n2 + ca
		node := &ListNode{
			Val:  n % 10,
			Next: nil,
		}
		cu.Next = node
		cu = cu.Next
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
			Next: nil,
		}
		cu.Next = node1
		cu = cu.Next
	}
	return dummnyhead.Next
}
