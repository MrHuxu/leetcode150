struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> bool {
        if target == nums[0] {
            return true;
        }

        let rotate_idx =
            Self::find_rotate_idx(&nums, &mut 0, &mut (nums.len() as i32 - 1)).unwrap();
        let mut left = (if target > nums[0] { 0 } else { rotate_idx + 1 }) as i32;
        let mut right = (if target > nums[0] {
            rotate_idx
        } else {
            nums.len() - 1
        }) as i32;

        while left <= right {
            let mid = (left + right) / 2;
            match nums[mid as usize].cmp(&target) {
                Ordering::Equal => return true,
                Ordering::Less => left = mid + 1,
                Ordering::Greater => right = mid - 1,
            }
        }

        false
    }

    fn find_rotate_idx(nums: &Vec<i32>, left: &mut i32, right: &mut i32) -> Option<usize> {
        while left <= right {
            let mid = ((*left + *right) / 2) as usize;
            if mid == nums.len() - 1 || nums[mid] > nums[mid + 1] {
                return Some(mid);
            }

            match nums[mid].cmp(&nums[0]) {
                Ordering::Less => *right = mid as i32 - 1,
                Ordering::Greater => *left = mid as i32 + 1,
                Ordering::Equal => {
                    if let Some(val) = Self::find_rotate_idx(nums, left, &mut (mid as i32 - 1)) {
                        return Some(val);
                    } else {
                        return Self::find_rotate_idx(nums, &mut (mid as i32 + 1), right);
                    }
                }
            }
        }
        None
    }
}

#[test]
fn test() {
    assert_eq!(Solution::search(vec![2, 5, 6, 0, 0, 1, 2], 0), true);
    assert_eq!(
        Solution::search(vec![4, 5, 6, 6, 7, 0, 1, 2, 4, 4], 0),
        true
    );
    assert_eq!(Solution::search(vec![2, 5, 6, 0, 0, 1, 2], 3), false);
    assert_eq!(
        Solution::search(
            vec![1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1],
            2
        ),
        true
    );
}
