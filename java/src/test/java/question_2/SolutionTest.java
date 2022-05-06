package question_2;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

import util.ListNode;

public class SolutionTest {
    @Test
    void addTwoNumbers() {
        assertThat(
                new Solution().addTwoNumbers(
                        ListNode.buildByArray(new int[] { 2, 4, 3 }),
                        ListNode.buildByArray(new int[] { 5, 6, 4 })))
                .usingRecursiveComparison()
                .isEqualTo(ListNode.buildByArray(new int[] { 7, 0, 8 }));

        assertThat(new Solution().addTwoNumbers(
                ListNode.buildByArray(new int[] { 0 }),
                ListNode.buildByArray(new int[] { 0 })))
                .usingRecursiveComparison()
                .isEqualTo(ListNode.buildByArray(new int[] { 0 }));

        assertThat(
                new Solution().addTwoNumbers(
                        ListNode.buildByArray(new int[] { 9, 9, 9, 9, 9, 9, 9 }),
                        ListNode.buildByArray(new int[] { 9, 9, 9, 9 })))
                .usingRecursiveComparison()
                .isEqualTo(ListNode.buildByArray(new int[] { 8, 9, 9, 9, 0, 0, 0, 1 }));
    }
}
