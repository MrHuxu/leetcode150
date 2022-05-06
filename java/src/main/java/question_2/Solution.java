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
        ListNode dummyHead = new ListNode();

        ListNode tmp = dummyHead;
        int carry = 0;
        while (l1 != null || l2 != null) {
            if (l1 != null && l2 != null) {
                carry += l1.val + l2.val;
                l1 = l1.next;
                l2 = l2.next;
            } else if (l1 != null) {
                carry += l1.val;
                l1 = l1.next;
            } else if (l2 != null) {
                carry += l2.val;
                l2 = l2.next;
            }

            tmp.next = new ListNode(carry % 10);
            tmp = tmp.next;
            carry /= 10;
        }
        if (carry != 0) {
            tmp.next = new ListNode(carry);
        }

        return dummyHead.next;
    }
}