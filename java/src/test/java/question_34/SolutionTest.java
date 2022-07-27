package question_34;

import static org.assertj.core.api.Assertions.assertThat;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void searchRange() {
        assertThat(new Solution().searchRange(new int[] { 1, 2, 3 }, 2)).usingDefaultComparator()
                .isEqualTo(new int[] { 1, 1 });
        assertThat(new Solution().searchRange(new int[] { 5, 7, 7, 8, 8, 10 }, 8)).usingDefaultComparator()
                .isEqualTo(new int[] { 3, 4 });
        assertThat(new Solution().searchRange(new int[] { 5, 7, 7, 8, 8, 10 }, 6)).usingDefaultComparator()
                .isEqualTo(new int[] { -1, -1 });
        assertThat(new Solution().searchRange(new int[] {}, 1)).usingDefaultComparator()
                .isEqualTo(new int[] { -1, -1 });
    }
}
