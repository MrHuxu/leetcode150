package question_2;

import util.ListNode;

/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        return addTwoNumbers(l1, l2, 0);
    }

    private ListNode addTwoNumbers(ListNode l1, ListNode l2, int carry) {
        if (l1 != null && l2 != null) {
            carry += l1.val + l2.val;
            return new ListNode(carry % 10, addTwoNumbers(l1.next, l2.next, carry / 10));
        } else if (l1 != null) {
            carry += l1.val;
            return new ListNode(carry % 10, addTwoNumbers(l1.next, l2, carry / 10));
        } else if (l2 != null) {
            carry += l2.val;
            return new ListNode(carry % 10, addTwoNumbers(l1, l2.next, carry / 10));
        } else {
            if (carry == 0) {
                return null;
            } else {
                return new ListNode(carry);
            }
        }
    }
}