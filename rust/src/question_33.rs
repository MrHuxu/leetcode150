struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut left: i32 = 0;
        let mut right: i32 = nums.len() as i32 - 1;

        while left <= right {
            let mid = ((left + right) / 2) as usize;
            match nums[mid].cmp(&nums[0]) {
                Ordering::Less => right = mid as i32 - 1,
                _ => left = mid as i32 + 1,
            };
        }
        let rotate_idx = (left + right) / 2;
        if target >= nums[0] {
            left = 0;
            right = rotate_idx;
        } else {
            left = rotate_idx + 1;
            right = nums.len() as i32 - 1;
        }

        while left <= right {
            let mid = (left + right) / 2;
            match nums[mid as usize].cmp(&target) {
                Ordering::Less => left = mid + 1,
                Ordering::Greater => right = mid - 1,
                Ordering::Equal => return mid as i32,
            };
        }

        -1
    }
}

#[test]
fn test() {
    assert_eq!(Solution::search(vec![8, 9, 2, 3, 4], 9), 1);
    assert_eq!(Solution::search(vec![1, 3], 0), -1);
    assert_eq!(Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 0), 4);
    assert_eq!(Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 3), -1);
    assert_eq!(Solution::search(vec![1], 0), -1);
    assert_eq!(Solution::search(vec![1, 3], 3), 1);
    assert_eq!(Solution::search(vec![3, 5, 1], 1), 2);
    assert_eq!(Solution::search(vec![5, 1, 3], 1), 1);
}
