package question_7;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void reverse() {
        assertEquals(321, new Solution().reverse(123));
        assertEquals(0, new Solution().reverse(1534236469));
    }
}
