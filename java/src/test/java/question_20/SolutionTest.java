package question_20;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void isValid() {
        assertTrue(new Solution().isValid("()"));
        assertTrue(new Solution().isValid("()[]{}"));
        assertFalse(new Solution().isValid("(]"));
    }
}