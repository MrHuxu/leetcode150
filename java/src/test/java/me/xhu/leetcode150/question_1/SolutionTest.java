package me.xhu.leetcode150.question_1;

import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void testTwoSum() {
        assertArrayEquals(new int[] { 0, 1 }, new Solution().twoSum(new int[] { 2, 7, 11, 15 }, 9));
        assertArrayEquals(new int[] { 1, 2 }, new Solution().twoSum(new int[] { 3, 2, 4 }, 6));
    }
}
