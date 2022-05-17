package question_3;

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        int ret = 0;

        int front = 0;
        Map<Character, Integer> m = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            char ch = s.charAt(i);
            if (m.getOrDefault(ch, 0) == 0) {
                if (i - front + 1 > ret) {
                    ret = i - front + 1;
                }
            } else {
                while (m.get(ch) != 0) {
                    m.computeIfPresent(s.charAt(front), (key, val) -> val - 1);
                    front++;
                }
            }
            m.put(ch, 1);
        }

        return ret;
    }
}
