import unittest
from typing import List, Optional

from list import ListNode, buildList, compareList

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        if not lists:
            return None

        while len(lists) > 1:
            next: List[Optional[ListNode]] = []
            for i in range(0, len(lists), 2):
                if i + 1 == len(lists):
                    next.append(lists[i])
                else:
                    next.append(self.mergeTwoLists(lists[i], lists[i + 1]))
            lists = next

        return lists[0]

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
        self.assertTrue(compareList(solution.mergeKLists([
            buildList([1, 4, 5]),
            buildList([1, 3, 4]),
            buildList([2, 6])
        ]), buildList([1, 1, 2, 3, 4, 4, 5, 6])))
