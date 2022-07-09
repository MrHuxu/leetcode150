package question_97;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void isInterleave() {
        assertTrue(new Solution().isInterleave("a", "b", "ba"));
        assertTrue(new Solution().isInterleave("aabcc", "dbbca", "aadbbcbcac"));
        assertFalse(new Solution().isInterleave("aabcc", "dbbca", "aadbbbaccc"));
        assertFalse(new Solution().isInterleave("", "", "a"));
    }
}
