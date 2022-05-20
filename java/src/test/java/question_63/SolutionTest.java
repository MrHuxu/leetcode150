package question_63;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void uniquePathsWithObstacles() {
        assertEquals(2, new Solution().uniquePathsWithObstacles(new int[][] {
                new int[] { 0, 0, 0 }, new int[] { 0, 1, 0 }, new int[] { 0, 0, 0 }
        }));
        assertEquals(1, new Solution().uniquePathsWithObstacles(new int[][] {
                new int[] { 0, 1 }, new int[] { 0, 0 }
        }));
        assertEquals(0, new Solution().uniquePathsWithObstacles(new int[][] {
                new int[] { 0, 0 }, new int[] { 0, 1 }
        }));
    }
}
