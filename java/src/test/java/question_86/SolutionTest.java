package question_86;

import static org.junit.jupiter.api.Assertions.assertEquals;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.ListNode;

public class SolutionTest {
    @Test
    void partition() {
        assertThat(new Solution().partition(ListNode.buildByArray(new int[] { 1, 4, 3, 2, 5, 2 }), 3))
                .usingRecursiveComparison().isEqualTo(ListNode.buildByArray(new int[] { 1, 2, 2, 4, 3, 5 }));
    }
}
