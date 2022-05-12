package question_47;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void permuteUnique() {
        assertEquals(Arrays.asList(
                Arrays.asList(1, 1, 2),
                Arrays.asList(1, 2, 1),
                Arrays.asList(2, 1, 1)), new Solution().permuteUnique(new int[] { 1, 1, 2 }));

        assertEquals(Arrays.asList(
                Arrays.asList(1, 2, 3),
                Arrays.asList(1, 3, 2),
                Arrays.asList(2, 1, 3),
                Arrays.asList(2, 3, 1),
                Arrays.asList(3, 1, 2),
                Arrays.asList(3, 2, 1)), new Solution().permuteUnique(new int[] { 1, 2, 3 }));
    }
}
