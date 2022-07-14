package question_105;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.TreeNode;

public class SolutionTest {
    @Test
    void buildTree() {
        assertThat(new Solution().buildTree(new int[] { 3, 9, 20, 15, 7 }, new int[] { 9, 3, 15, 20, 7 }))
                .usingRecursiveComparison()
                .isEqualTo(TreeNode.buildByArray(new Integer[] { 3, 9, 20, null, null, 15, 7 }));
    }
}
