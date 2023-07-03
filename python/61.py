from typing import Optional
import unittest
from list import ListNode, buildList, compareList


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return head

        length, tmp = 0, head
        while tmp:
            length += 1
            tmp = tmp.next
        k = k % length

        dummy_head = ListNode(0, head)
        fast, slow = dummy_head, dummy_head
        for _ in range(k):
            fast = fast.next if fast else None
        while fast and fast.next:
            fast = fast.next
            slow = slow.next if slow else None
        if fast:
            fast.next = head
        dummy_head.next = slow.next if slow else None
        if slow:
            slow.next = None

        return dummy_head.next


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.rotateRight(
            buildList([1, 2, 3, 4, 5]), 2), buildList([4, 5, 1, 2, 3])))
