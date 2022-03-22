struct Solution;

use std::cmp::max;

impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut max_distance = 0;
        for (i, num) in nums.iter().enumerate() {
            if i > max_distance {
                return false;
            }

            max_distance = max(max_distance, i + *num as usize);
        }

        true
    }
}

#[test]
fn test() {
    assert_eq!(Solution::can_jump(vec![2, 3, 1, 1, 4]), true);
    assert_eq!(Solution::can_jump(vec![3, 2, 1, 0, 4]), false);
}
