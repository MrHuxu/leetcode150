struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        if target <= nums[0] {
            return 0;
        }
        if target > *nums.last().unwrap() {
            return nums.len() as i32;
        }

        let mut left = 0;
        let mut right = nums.len() - 1;
        while left < right {
            let mid = (left + right) / 2;
            match nums[mid].cmp(&target) {
                Ordering::Less => left = mid + 1,
                _ => right = mid,
            }
        }
        ((left + right) / 2) as i32
    }
}

#[test]
fn test() {
    assert_eq!(Solution::search_insert(vec![1, 3, 5, 6], 5), 2);
    assert_eq!(Solution::search_insert(vec![1, 3, 5, 6], 2), 1);
    assert_eq!(Solution::search_insert(vec![1, 3, 5, 6], 1), 0);
    assert_eq!(Solution::search_insert(vec![1, 3, 5, 6], 7), 4);
}
