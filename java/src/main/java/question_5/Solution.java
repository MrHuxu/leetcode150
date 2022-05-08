package question_4;

class Solution {
    public String longestPalindrome(String s) {
        if (s.length() == 0) {
            return "";
        }

        int resultLeft = 0;
        int resultRight = 0;

        for (int i = 0; i < s.length(); i++) {
            for (int j = s.length() - 1; j >= i + (resultRight - resultLeft); j--) {
                if (isPalindrome(i, j, s)) {
                    resultLeft = i;
                    resultRight = j;
                    break;
                }
            }
        }

        return s.substring(resultLeft, resultRight + 1);
    }

    private boolean isPalindrome(int i, int j, String s) {
        while (i <= j) {
            if (s.charAt(i) != s.charAt(j)) {
                return false;
            }

            i++;
            j--;
        }
        return true;
    }
}