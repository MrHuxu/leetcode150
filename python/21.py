from typing import Optional
import unittest

from list import ListNode, buildList, compareList


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if not list1 and not list2:
            return None
        if not list1:
            return list2
        if not list2:
            return list1
        return ListNode(
            list1.val if list1.val < list2.val else list2.val,
            self.mergeTwoLists(
                list1.next if list1.val < list2.val else list1,
                list2.next if list1.val >= list2.val else list2
            )
        )


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareList(solution.mergeTwoLists(
            buildList([1, 2, 4]),
            buildList([1, 3, 4])
        ), buildList([1, 1, 2, 3, 4, 4])))
