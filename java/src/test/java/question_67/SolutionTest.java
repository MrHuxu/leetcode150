package question_67;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void uniquePaths() {
        assertEquals("100", new Solution().addBinary("11", "1"));
        assertEquals("1", new Solution().addBinary("0", "1"));
        assertEquals("10101", new Solution().addBinary("1010", "1011"));
    }
}
