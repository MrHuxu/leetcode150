package question_17;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Map;

class Solution {
    public List<String> letterCombinations(String digits) {
        if (digits.length() == 0) {
            return new ArrayList<>();
        }

        Map<Character, Character[]> mapDigitToLetters = Map.of(
                '2', new Character[] { 'a', 'b', 'c' },
                '3', new Character[] { 'd', 'e', 'f' },
                '4', new Character[] { 'g', 'h', 'i' },
                '5', new Character[] { 'j', 'k', 'l' },
                '6', new Character[] { 'm', 'n', 'o' },
                '7', new Character[] { 'p', 'q', 'r', 's' },
                '8', new Character[] { 't', 'u', 'v' },
                '9', new Character[] { 'w', 'x', 'y', 'z' });

        return helper(0, mapDigitToLetters, digits);
    }

    private List<String> helper(int idx, Map<Character, Character[]> mapDigitToLetters, String digits) {
        if (idx == digits.length()) {
            return Arrays.asList("");
        }

        List<String> ret = new ArrayList<>();

        List<String> rest = helper(idx + 1, mapDigitToLetters, digits);
        for (String string : rest) {
            for (Character ch : mapDigitToLetters.get(digits.charAt(idx))) {
                StringBuilder sb = new StringBuilder();
                sb.append(ch);
                sb.append(string);
                ret.add(sb.toString());
            }
        }

        return ret;
    }
}