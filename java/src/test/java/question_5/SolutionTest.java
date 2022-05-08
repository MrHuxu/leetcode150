package question_4;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void longestPalindrome() {
        assertEquals("aba", new Solution().longestPalindrome("babad"));
        assertEquals("bb", new Solution().longestPalindrome("cbbd"));
    }
}
