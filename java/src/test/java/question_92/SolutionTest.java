package question_92;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.ListNode;

public class SolutionTest {
    @Test
    void reverseBetween() {
        assertThat(new Solution().reverseBetween(ListNode.buildByArray(new int[] { 1, 2, 3, 4, 5 }), 2, 4))
                .usingRecursiveComparison().isEqualTo(ListNode.buildByArray(new int[] { 1, 4, 3, 2, 5 }));
    }
}