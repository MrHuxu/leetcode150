package question_34;

class Solution {
    public int[] searchRange(int[] nums, int target) {
        int[] ret = new int[]{-1, -1};

        if (nums.length < 1) {
            return ret;
        }

        int left = 0;
        int right = nums.length - 1;
        while (left < right) {
            int mid = (left + right) / 2;
            if (nums[mid] >= target) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        if (nums[(left + right) / 2] == target) {
            ret[0] = (left + right) / 2;
        }

        if (ret[0] == -1) {
            return ret;
        }

        left = ret[0];
        right = nums.length -1;
        while (left <= right) {
            int mid = (left + right) / 2;
            if (nums[mid] > target) {
                right = mid - 1;
            } else  {
                left = mid + 1;
            }
        }
        ret[1] = (left + right) / 2;

        return ret;
    }
}