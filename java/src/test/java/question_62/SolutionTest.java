package question_62;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void uniquePaths() {
        assertEquals(28, new Solution().uniquePaths(3, 7));
        assertEquals(3, new Solution().uniquePaths(3, 2));
    }
}
