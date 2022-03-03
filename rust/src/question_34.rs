struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut ret = vec![-1, -1];

        if nums.len() == 0 {
            return ret;
        }

        let mut left: i32 = 0;
        let mut right: i32 = nums.len() as i32 - 1;
        while left < right {
            let mid = (left + right) / 2;
            match nums[mid as usize].cmp(&target) {
                Ordering::Less => left = mid + 1,
                Ordering::Equal | Ordering::Greater => right = mid,
            }
        }
        if nums[((left + right) / 2) as usize] == target {
            ret[0] = (left + right) / 2;
        }

        left = 0;
        right = nums.len() as i32 - 1;
        while left <= right {
            let mid = (left + right) / 2;
            match nums[mid as usize].cmp(&target) {
                Ordering::Less | Ordering::Equal => left = mid + 1,
                Ordering::Greater => right = mid - 1,
            }
        }
        if nums[((left + right) / 2) as usize] == target {
            ret[1] = (left + right) / 2;
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::search_range(vec![5, 7, 7, 8, 8, 10], 8),
        vec![3, 4]
    );
    assert_eq!(
        Solution::search_range(vec![5, 7, 7, 8, 8, 10], 6),
        vec![-1, -1]
    );
    assert_eq!(Solution::search_range(vec![], 0), vec![-1, -1]);
    assert_eq!(Solution::search_range(vec![1], 0), vec![-1, -1]);
    assert_eq!(Solution::search_range(vec![1, 2, 3], 2), vec![1, 1]);
}
