package question_3;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void lengthOfLongestSubstring() {
        assertEquals(3, new Solution().lengthOfLongestSubstring("abcabcbb"));
        assertEquals(1, new Solution().lengthOfLongestSubstring("bbbbb"));
        assertEquals(3, new Solution().lengthOfLongestSubstring("pwwkew"));
    }
}
