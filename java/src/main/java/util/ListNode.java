package util;

public class ListNode {
    public int val;
    public ListNode next;

    public ListNode() {
    }

    public ListNode(int val) {
        this.val = val;
    }

    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    public final static ListNode buildByArray(int[] vals) {
        ListNode dummyHead = new ListNode();

        ListNode tmp = dummyHead;
        for (int val : vals) {
            tmp.next = new ListNode(val);
            tmp = tmp.next;
        }

        return dummyHead.next;
    }
}