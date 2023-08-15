import unittest
from typing import Optional

from list import ListNode, buildList, compareList

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        smaller_head, greater_head = ListNode(0, head), ListNode(0)

        tmp, smaller, greater = head, smaller_head, greater_head
        while tmp:
            if tmp.val < x:
                smaller.next = tmp
                smaller = smaller.next
            else:
                greater.next = tmp
                greater = greater.next
            tmp = tmp.next
        smaller.next = greater_head.next
        greater.next = None

        return smaller_head.next


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.partition(
            buildList([1, 4, 3, 2, 5, 2]), 3), buildList([1, 2, 2, 4, 3, 5])))
