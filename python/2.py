import unittest
from typing import Optional

from list import ListNode, buildList, compareList

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        return self.helper(l1, l2, 0)

    def helper(self, l1: Optional[ListNode], l2: Optional[ListNode], carry: int) -> Optional[ListNode]:
        if l1 is None and l2 is None and carry == 0:
            return None
        val = carry + (0 if not l1 else l1.val) + (0 if not l2 else l2.val)
        return ListNode(val % 10, self.helper(l1.next if l1 else None, l2.next if l2 else None, val // 10))


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.addTwoNumbers(
            buildList([2, 4, 3]), buildList([5, 6, 4])), buildList([7, 0, 8])))
