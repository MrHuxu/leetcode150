package question_17;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void letterCombinations() {
        assertEquals(Arrays.asList("ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"),
                new Solution().letterCombinations("23"));
    }
}
