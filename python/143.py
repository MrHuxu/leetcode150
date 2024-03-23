from typing import Optional
from list import ListNode

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if not head:
            return None

        slow, fast = ListNode(next=head), head
        while True:
            slow, fast = slow.next, fast.next
            if not fast or not fast.next:
                break
            fast = fast.next

        reversed, iter, slow.next = None, slow.next, None
        while iter:
            tmp, iter = iter, iter.next
            tmp.next, reversed = reversed, tmp

        dummy_head, idx = ListNode(), 0
        while head or reversed:
            if idx % 2 == 0:
                dummy_head.next = head
                head = head.next
            else:
                dummy_head.next = reversed
                reversed = reversed.next
            idx += 1
            dummy_head = dummy_head.next
            dummy_head.next = None
