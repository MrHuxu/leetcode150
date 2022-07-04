package question_135;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void candy() {
        assertEquals(5, new Solution().candy(new int[] { 1, 0, 2 }));
        assertEquals(4, new Solution().candy(new int[] { 1, 2, 2 }));
        assertEquals(13, new Solution().candy(new int[] { 1, 2, 87, 87, 87, 2, 1 }));
    }
}
