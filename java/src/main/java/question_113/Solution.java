package question_113;

import java.util.ArrayList;
import java.util.Arrays;
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
    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        if (root == null) {
            return new ArrayList<>();
        }

        if (root.left == null && root.right == null && root.val == targetSum) {
            return new ArrayList<>(Arrays.asList(
                new ArrayList<>(Arrays.asList(root.val))));
        }

        List<List<Integer>> ret = new ArrayList<>();
        for (List<Integer> tmp : pathSum(root.left, targetSum - root.val)) {
            tmp.add(0, root.val);
            ret.add(tmp);
        }
        for (List<Integer> tmp : pathSum(root.right, targetSum - root.val)) {
            tmp.add(0, root.val);
            ret.add(tmp);
        }

        return ret;
    }
}