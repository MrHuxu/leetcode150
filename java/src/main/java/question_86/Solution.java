package question_86;

import util.ListNode;

class Solution {
    public ListNode partition(ListNode head, int x) {
        ListNode dummyHead = new ListNode();

        ListNode smallerTail = null;
        ListNode greaterHead = new ListNode();
        ListNode greaterTail = null;
        while (head != null) {
            ListNode curr = head;
            head = head.next;
            if (curr.val < x) {
                if (smallerTail == null) {
                    smallerTail = curr;
                    dummyHead.next = smallerTail;
                } else {
                    smallerTail.next = curr;
                    smallerTail = smallerTail.next;
                }
            } else {
                if (greaterTail == null) {
                    greaterTail = curr;
                    greaterHead.next = greaterTail;
                } else {
                    greaterTail.next = curr;
                    greaterTail = greaterTail.next;
                }
            }
        }
        (smallerTail == null ? dummyHead : smallerTail).next = greaterHead.next;
        if (greaterTail != null) {
            greaterTail.next = null;
        }

        return dummyHead.next;
    }
}
