# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        return self.helper(0, root)
    
    def helper(self, pre: int, node: Optional[TreeNode]) -> int:
        if not node:
            return 0
        curr = pre * 10 + node.val
        if not node.left and not node.right:
            return curr
        return self.helper(curr, node.left) + self.helper(curr, node.right)
