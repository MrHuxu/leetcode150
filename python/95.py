import unittest
from typing import List, Optional

from tree import TreeNode, buildTree


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def generateTrees(self, n: int) -> List[Optional[TreeNode]]:
        return self.helper(1, n)

    def helper(self, start: int, end: int) -> List[Optional[TreeNode]]:
        if start > end:
            return [None]
        
        ret: List[Optional[TreeNode]] = []
        for i in range(start, end + 1):
            for left in self.helper(start, i - 1):
                for right in self.helper(i + 1, end):
                    ret.append(TreeNode(i, left, right))
        return ret
