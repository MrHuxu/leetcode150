package question_102;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.TreeNode;

public class SolutionTest {
    @Test
    void levelOrder() {
        assertThat(new Solution().levelOrder(TreeNode.buildByArray(new Integer[] { 3, 9, 20, null, null, 15, 7 })))
                .usingRecursiveComparison()
                .isEqualTo(List.of(List.of(3), List.of(9, 20), List.of(15, 7)));
    }
}
