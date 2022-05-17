package question_94;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Test;

import util.TreeNode;

public class SolutionTest {
    @Test
    void inorderTraversal() {
        assertEquals(Arrays.asList(1, 3, 2),
                new Solution().inorderTraversal(TreeNode.buildByArray(new Integer[] { 1, null, 2, 3 })));
    }
}
