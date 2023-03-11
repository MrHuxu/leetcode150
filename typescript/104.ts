import { buildTree, TreeNode } from './tree';

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

function maxDepth(root: TreeNode | null): number {
    return root ? Math.max(maxDepth(root.left), maxDepth(root.right)) + 1 : 0;
};

test('104', () => {
    expect(maxDepth(buildTree([3, 9, 20, null, null, 15, 7]))).toBe(3);
    expect(maxDepth(buildTree([1, null, 2]))).toBe(2);
});