package question_114;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.TreeNode;

public class SolutionTest {
    @Test
    void flatten() {
        TreeNode root = TreeNode.buildByArray(new Integer[] { 1, 2, 5, 3, 4, null, 6 });
        new Solution().flatten(root);
        assertThat(root)
                .usingRecursiveComparison()
                .isEqualTo(TreeNode.buildByArray(new Integer[] { 1, null, 2, null, 3, null, 4, null, 5, null, 6 }));
    }
}