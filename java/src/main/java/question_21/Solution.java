package question_21;

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
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        if (list1 == null && list2 == null) {
            return null;
        } else if (list1 == null) {
            return list2;
        } else if (list2 == null) {
            return list1;
        } else {
            ListNode node;
            if (list1.val < list2.val) {
                node = list1;
                node.next = mergeTwoLists(list1.next, list2);
            } else {
                node = list2;
                node.next = mergeTwoLists(list1, list2.next);
            }
            return node;
        }
    }
}