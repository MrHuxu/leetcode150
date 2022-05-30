package question_29;

class Solution {
    public int divide(int dividend, int divisor) {
        int sym = 1;
        long dividendL = Math.abs((long) (dividend));
        long divisorL = Math.abs((long) (divisor));
        if (dividend < 0)
            sym *= -1;
        if (divisor < 0)
            sym *= -1;

        long ret = 0;

        while (dividendL >= divisorL) {
            long i = 1;
            while (divisorL * i * 2 <= dividendL)
                i *= 2;

            ret += i;
            dividendL -= divisorL * i;
        }

        if (ret * sym > Integer.MAX_VALUE)
            return Integer.MAX_VALUE;
        else if (ret * sym < Integer.MIN_VALUE)
            return Integer.MIN_VALUE;

        return (int) (ret * sym);
    }
}