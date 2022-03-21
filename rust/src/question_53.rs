struct Solution;

use std::cmp::max;

impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut ret = nums[0];
        let mut pre = nums[0];

        for i in 1..nums.len() {
            pre = if pre < 0 { nums[i] } else { pre + nums[i] };
            ret = max(ret, pre);
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::max_sub_array(vec![-2, 1, -3, 4, -1, 2, 1, -5, 4]),
        6
    );
    assert_eq!(Solution::max_sub_array(vec![1]), 1);
}
