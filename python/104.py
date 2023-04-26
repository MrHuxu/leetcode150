from typing import Optional
import unittest

from tree import TreeNode, buildTree

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))


class TestSolution(unittest.TestCase):
    def test_maxDepth(self):
        solution = Solution()
        self.assertEqual(solution.maxDepth(
            buildTree([3, 9, 20, None, None, 15, 7])), 3)
