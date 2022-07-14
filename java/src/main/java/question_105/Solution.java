package question_105;

import java.util.Arrays;

import util.TreeNode;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        for (int i = 0; i < inorder.length; i++) {
            if (inorder[i] == preorder[0]) {
                return new TreeNode(
                        preorder[0],
                        buildTree(Arrays.copyOfRange(preorder, 1, i + 1), Arrays.copyOfRange(inorder, 0, i)),
                        buildTree(Arrays.copyOfRange(preorder, i + 1, preorder.length),
                                Arrays.copyOfRange(inorder, i + 1, inorder.length)));
            }
        }

        return null;
    }
}