struct Solution;

use std::cmp::min;

impl Solution {
    pub fn jump(nums: Vec<i32>) -> i32 {
        let mut dp = vec![0; nums.len()];

        for i in 0..nums.len() - 1 {
            if nums[i] == 0 {
                continue;
            }

            if i + nums[i] as usize >= nums.len() - 1 {
                return dp[i] + 1;
            }

            for j in i + 1..=min(nums.len() - 1, i + nums[i] as usize) {
                if dp[j] == 0 {
                    dp[j] = dp[i] + 1;
                }
            }
        }

        dp[nums.len() - 1]
    }
}

#[test]
fn test() {
    assert_eq!(Solution::jump(vec![1, 2, 3]), 2);
    assert_eq!(Solution::jump(vec![2, 3, 1, 1, 4]), 2);
    assert_eq!(Solution::jump(vec![2, 3, 0, 1, 4]), 2);
}
