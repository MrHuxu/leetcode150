package question_128;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void longestConsecutive() {
        assertEquals(4, new Solution().longestConsecutive(new int[] { 100, 4, 200, 1, 3, 2 }));
        assertEquals(9, new Solution().longestConsecutive(new int[] { 0, 3, 7, 2, 5, 8, 4, 6, 0, 1 }));
    }
}
