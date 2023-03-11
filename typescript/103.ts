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

function zigzagLevelOrder(root: TreeNode | null): number[][] {
    if (!root) return [];

    const res: number[][] = [];
    let level = [root];
    let fromLeftToRight = true;
    while (level.length) {
        const nextLevel: TreeNode[] = [];
        const vals: number[] = [];
        for (let node of level) {
            vals.push(node.val);
            if (node.left) nextLevel.push(node.left);
            if (node.right) nextLevel.push(node.right);
        }
        if (!fromLeftToRight) vals.reverse();
        res.push(vals);
        fromLeftToRight = !fromLeftToRight;
        level = nextLevel;
    }

    return res;
};

test('103', () => {
    expect(zigzagLevelOrder(buildTree([3, 9, 20, null, null, 15, 7]))).toStrictEqual([[3], [20, 9], [15, 7]]);
    expect(zigzagLevelOrder(buildTree([1]))).toStrictEqual([[1]]);
});