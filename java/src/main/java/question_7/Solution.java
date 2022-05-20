package question_7;

class Solution {
    public int reverse(int x) {
        return helper(0, x);
    }

    private int helper(int pre, int num) {
        if (num == 0)
            return (int) pre;

        if ((num > 0 && (pre * 10 + num % 10) / 10 < pre) || (num < 0 && (pre * 10 + num % 10) / 10 > pre))
            return 0;

        return helper(pre * 10 + num % 10, num / 10);
    }
}