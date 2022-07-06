package question_128;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

class Solution {
    public int longestConsecutive(int[] nums) {
        int ret = 0;

        Map<Integer, Integer> mapHeadToTail = new HashMap<>();
        Map<Integer, Integer> mapTailToHead = new HashMap<>();
        Set<Integer> existence = new HashSet<>();
        for (Integer num : nums) {
            if (existence.contains(num))
                continue;
            else
                existence.add(num);

            if (mapHeadToTail.containsKey(num + 1) && mapTailToHead.containsKey(num - 1)) {
                int tail = mapHeadToTail.get(num + 1);
                int head = mapTailToHead.get(num - 1);
                mapHeadToTail.put(head, tail);
                mapTailToHead.put(tail, head);
                ret = Math.max(ret, tail - head + 1);
                continue;
            }

            if (mapHeadToTail.containsKey(num + 1)) {
                int tail = mapHeadToTail.get(num + 1);
                mapHeadToTail.put(num, tail);
                mapTailToHead.put(tail, num);
                ret = Math.max(ret, tail - num + 1);
                continue;
            }

            if (mapTailToHead.containsKey(num - 1)) {
                int head = mapTailToHead.get(num - 1);
                mapHeadToTail.put(head, num);
                mapTailToHead.put(num, head);
                ret = Math.max(ret, num - head + 1);
                continue;
            }

            ret = Math.max(ret, 1);
            mapHeadToTail.put(num, num);
            mapTailToHead.put(num, num);
        }

        return ret;
    }
}