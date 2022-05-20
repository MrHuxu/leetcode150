package question_6;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }

        Map<Integer, ArrayList<Character>> m = new HashMap<>();
        boolean vertical = true;
        int line = 0;
        for (int i = 0; i < s.length(); i++) {
            m.putIfAbsent(line, new ArrayList<>());
            m.get(line).add(s.charAt(i));

            line += vertical ? 1 : -1;
            if (line == numRows - 1) {
                vertical = false;
            } else if (line == 0) {
                vertical = true;
            }
        }

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < numRows; i++) {
            if (!m.containsKey(i))
                continue;

            for (Character c : m.get(i))
                sb.append(c);
        }
        return sb.toString();
    }
}