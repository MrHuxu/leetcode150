package question_113;

import static org.assertj.core.api.Assertions.assertThat;

import java.util.List;

import org.junit.jupiter.api.Test;

import util.TreeNode;

public class SolutionTest {
    @Test
    void pathSum() {
        TreeNode root = TreeNode.buildByArray(new Integer[] { 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1 });
        List<List<Integer>> ret = new Solution().pathSum(root, 22);
        assertThat(ret)
                .usingRecursiveComparison()
                .isEqualTo(List.of(
                        List.of(5, 4, 11, 2), List.of(5, 8, 4, 5)));
    }
}