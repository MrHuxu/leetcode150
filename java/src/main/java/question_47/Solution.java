package question_47;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        Arrays.sort(nums);
        return helper(nums, new HashSet<>());
    }

    private List<List<Integer>> helper(int[] nums, Set<Integer> used) {
        List<List<Integer>> ret = new ArrayList<>();

        for (int i = 0; i < nums.length; i++) {
            if (used.contains(i) || (i > 0 && nums[i] == nums[i - 1] && !used.contains(i - 1))) {
                continue;
            }

            used.add(i);
            for (List<Integer> next : helper(nums, used)) {
                List<Integer> tmp = new ArrayList<>(Arrays.asList(nums[i]));
                tmp.addAll(next);
                ret.add(tmp);
            }
            used.remove(i);
        }

        if (ret.isEmpty()) {
            ret.add(new ArrayList<>());
        }
        return ret;
    }
}