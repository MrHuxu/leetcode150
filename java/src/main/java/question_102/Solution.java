package question_102;

import java.util.ArrayList;
import java.util.List;

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
    public List<List<Integer>> levelOrder(TreeNode root) {
        if (root == null) {
            return new ArrayList<>();
        }

        List<List<Integer>> ret = new ArrayList<>();

        List<TreeNode> level = List.of(root);
        while (!level.isEmpty()) {
            List<Integer> nums = new ArrayList<>();
            List<TreeNode> nextLevel = new ArrayList<>();
            for (TreeNode node : level) {
                nums.add(node.val);

                if (node.left != null)
                    nextLevel.add(node.left);
                if (node.right != null)
                    nextLevel.add(node.right);
            }
            ret.add(nums);
            level = nextLevel;
        }

        return ret;
    }
}