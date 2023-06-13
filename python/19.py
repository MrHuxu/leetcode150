import unittest
from typing import Optional

from list import ListNode, buildList, compareList

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummyHead = ListNode(0, head)

        fast, slow = dummyHead, dummyHead
        while n >= 0:
            fast = fast.next if fast else None
            n -= 1
        while fast:
            fast = fast.next if fast else None
            slow = slow.next if slow else None
        if slow:
            slow.next = slow.next.next if slow.next else None

        return dummyHead.next


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.removeNthFromEnd(
            buildList([1, 2, 3, 4, 5]), 2), buildList([1, 2, 3, 5])))
