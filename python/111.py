from typing import Optional

from tree import TreeNode

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        if not root.left and not root.right:
            return 1

        ret = -1
        if root.left:
            ret = 1 + self.minDepth(root.left)
        if root.right:
            ret = 1 + self.minDepth(root.right) if ret == -1 else min(ret, 1 + self.minDepth(root.right))
        return ret
            