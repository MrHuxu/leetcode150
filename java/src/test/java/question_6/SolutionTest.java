package question_6;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
    @Test
    void convert() {
        assertEquals("PAHNAPLSIIGYIR", new Solution().convert("PAYPALISHIRING", 3));
        assertEquals("PINALSIGYAHRPI", new Solution().convert("PAYPALISHIRING", 4));
    }
}
