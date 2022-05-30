package question_29;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void divide() {
        assertEquals(3, new Solution().divide(10, 3));
        assertEquals(-2, new Solution().divide(7, -3));
        assertEquals(1, new Solution().divide(1, 1));
        assertEquals(2147483647, new Solution().divide(-2147483648, -1));
    }
}
