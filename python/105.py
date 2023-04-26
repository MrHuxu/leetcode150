import unittest
from typing import List, Optional

from tree import TreeNode, buildTree, compareTree

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        for i in range(len(inorder)):
            if inorder[i] == preorder[0]:
                return TreeNode(inorder[i],
                                self.buildTree(preorder[1:i + 1], inorder[:i]),
                                self.buildTree(preorder[i + 1:], inorder[i + 1:]))
        return None


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(compareTree(solution.buildTree([3, 9, 20, 15, 7], [
                        9, 3, 15, 20, 7]), buildTree([3, 9, 20, None, None, 15, 7])))
