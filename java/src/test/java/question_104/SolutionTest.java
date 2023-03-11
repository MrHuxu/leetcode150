package question_104;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

import util.TreeNode;

public class SolutionTest {
    @Test
    void maxDepth() {
        assertEquals(3, new Solution().maxDepth(TreeNode.buildByArray(new Integer[] { 3, 9, 20, null, null, 15, 7 })));
        assertEquals(2, new Solution().maxDepth(TreeNode.buildByArray(new Integer[] { 1, null, 2 })));
    }
}