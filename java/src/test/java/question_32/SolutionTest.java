package question_32;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void longestValidParentheses() {
        assertEquals(2, new Solution().longestValidParentheses("(()"));
        assertEquals(4, new Solution().longestValidParentheses(")()())"));
        assertEquals(0, new Solution().longestValidParentheses(""));
    }
}
