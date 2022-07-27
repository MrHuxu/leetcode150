package question_114;

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
    public void flatten(TreeNode root) {
        helper(root);
    }

    private TreeNode helper(TreeNode root) {
        if (root == null) {
            return null;
        }

        if (root.left != null) {
            TreeNode tail = helper(root.left);
            tail.right = root.right;
            root.right = root.left;
            root.left = null;
        }

        if (root.right != null) {
            return helper(root.right);
        }

        return root;
    }
}