import unittest
from typing import Optional

from list import ListNode, buildList, compareList

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        dummy_head = ListNode(0)

        iter = dummy_head
        for _ in range(left - 1):
            iter.next = head
            iter = iter.next
            head = head.next

        reversed_head, reversed_tail = None, None
        for _ in range(right - left + 1):
            if not reversed_tail:
                reversed_tail = head

            tmp = head
            head = head.next
            tmp.next = reversed_head
            reversed_head = tmp

        iter.next = reversed_head
        reversed_tail.next = head

        return dummy_head.next


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.reverseBetween(
            buildList([1, 2, 3, 4, 5]), 2, 4), buildList([1, 4, 3, 2, 5])))
