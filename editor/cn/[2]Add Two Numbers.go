package main

// [2] Add Two Numbers ( add-two-numbers )
//You are given two non-empty linked lists representing two non-negative
//integers. The digits are stored in reverse order, and each of their nodes contains a
//single digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the
//number 0 itself.
//
//
// Example 1:
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//
//
// Example 2:
//
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//
//
// Example 3:
//
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//
// Constraints:
//
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have
//leading zeros.
//
//
// Related Topics 递归 链表 数学 👍 9951 👎 0

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head, current       *ListNode
		nextVal, currentVal int
	)

	for !(l1 == nil && l2 == nil) {
		if l1 != nil && l2 != nil {
			currentVal = (nextVal + l1.Val + l2.Val) % 10
			nextVal = (nextVal + l1.Val + l2.Val) / 10

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			currentVal = (nextVal + l1.Val) % 10
			nextVal = (nextVal + l1.Val) / 10

			l1 = l1.Next
		} else if l2 != nil && l1 == nil {
			currentVal = (nextVal + l2.Val) % 10
			nextVal = (nextVal + l2.Val) / 10

			l2 = l2.Next
		}

		if head == nil {
			head = &ListNode{Val: currentVal}
			current = head
		} else {
			current.Next = &ListNode{Val: currentVal}
			current = current.Next
		}
	}

	if nextVal > 0 {
		current.Next = &ListNode{Val: nextVal}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
