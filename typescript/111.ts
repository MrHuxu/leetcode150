/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function minDepth(root: TreeNode | null): number {
    if (!root) return 0;
    if (!root.left && !root.right) return 1;

    let ret = -1;
    if (root.left) ret = 1 + minDepth(root.left);
    if (root.right) ret = ret === -1 ? 1 + minDepth(root.right) : Math.min(ret, 1 + minDepth(root.right));
    return ret;
};