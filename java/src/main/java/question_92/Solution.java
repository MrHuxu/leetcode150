package question_92;

import util.ListNode;

class Solution {
    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode dummyHead = new ListNode();

        ListNode iter = dummyHead;
        for (int i = 1; i < left; i++) {
            iter.next = head;
            iter = iter.next;
            head = head.next;
        }

        ListNode reversedHead = null;
        ListNode reversedTail = null;
        for (int i = left; i <= right; i++) {
            if (reversedTail == null) {
                reversedTail = head;
            }

            ListNode tmp = head;
            head = head.next;
            tmp.next = reversedHead;
            reversedHead = tmp;
        }

        iter.next = reversedHead;
        reversedTail.next = head;

        return dummyHead.next;
    }
}