package question_1;

import java.util.HashMap;
import java.util.Map;

// code
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<Integer, Integer>();
        int ret[] = new int[2];

        for (int i = 0; i < nums.length; i++) {
            if (m.containsKey(target - nums[i])) {
                ret[0] = m.get(target - nums[i]);
                ret[1] = i;
                break;
            }
            m.put(nums[i], i);
        }

        return ret;
    }
}
