from typing import List, Optional
from list import ListNode
import unittest


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if not head or not head.next:
            return False

        slow, fast = head, head.next
        while fast and fast.next:
            if fast.next.next:
                fast = fast.next.next
            else:
                return False
            if slow == fast:
                return True
            if slow:
                slow = slow.next

        return False


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        node1, node2, node3, node4 = ListNode(
            3), ListNode(2), ListNode(0), ListNode(-4)
        node1.next = node2
        node2.next = node3
        node3.next = node4
        node4.next = node2
        self.assertTrue(solution.hasCycle(node1))
