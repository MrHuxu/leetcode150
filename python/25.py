import unittest
from typing import Optional

from list import ListNode, buildList, compareList

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummyHead = ListNode(0, head)
        self.helper(dummyHead, head, head, None, k, 0)
        return dummyHead.next

    def helper(self, head: Optional[ListNode], start: Optional[ListNode], curr: Optional[ListNode], reversed: Optional[ListNode], k: int, reversedLen: int):
        if reversedLen == k:
            if head:
                head.next = reversed
            for _ in range(reversedLen):
                head = head.next if head else None
            self.helper(head, curr, curr, None, k, 0)
        elif not curr:
            if head:
                head.next = start
        else:
            self.helper(head, start, curr.next, ListNode(
                curr.val, reversed), k, reversedLen + 1)


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.reverseKGroup(
            buildList([1, 2, 3, 4, 5]), 2), buildList([2, 1, 4, 3, 5])))
