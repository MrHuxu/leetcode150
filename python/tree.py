from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def buildTree(vals: List[int | None]) -> Optional[TreeNode]:
    if not vals or vals[0] is None:
        return None
    root = TreeNode(vals[0])
    nodes = [root]
    vals = vals[1:]
    while vals:
        node = nodes.pop(0)
        if vals[0] is not None:
            node.left = TreeNode(vals[0])
            nodes.append(node.left)
        if vals[1] is not None:
            node.right = TreeNode(vals[1])
            nodes.append(node.right)
        vals = vals[2:]
    return root


def compareTree(t1: Optional[TreeNode], t2: Optional[TreeNode]) -> bool:
    if not t1 and not t2:
        return True
    if not t1 or not t2:
        return False
    if t1.val != t2.val:
        return False
    return compareTree(t1.left, t2.left) and compareTree(t1.right, t2.right)
