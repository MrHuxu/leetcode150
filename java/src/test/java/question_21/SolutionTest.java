package question_21;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.ListNode;

public class SolutionTest {
    @Test
    void mergeTwoLists() {
        assertThat(new Solution().mergeTwoLists(
                ListNode.buildByArray(new int[] { 1, 2, 4 }),
                ListNode.buildByArray(new int[] { 1, 3, 4 })))
                .usingRecursiveComparison()
                .isEqualTo(ListNode.buildByArray(new int[] { 1, 1, 2, 3, 4, 4 }));

        assertThat(new Solution().mergeTwoLists(null, null))
                .usingRecursiveComparison()
                .isEqualTo(null);

        assertThat(new Solution().mergeTwoLists(
                null, ListNode.buildByArray(new int[] { 0 })))
                .usingRecursiveComparison()
                .isEqualTo(ListNode.buildByArray(new int[] { 0 }));
    }
}
