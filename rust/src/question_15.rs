struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut nums = nums;
        nums.sort();

        let mut ret: Vec<Vec<i32>> = Vec::new();
        for i in 0..nums.len() {
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }

            let mut left = i + 1;
            let mut right = nums.len() - 1;
            while left < right {
                match (nums[i] + nums[left] + nums[right]).cmp(&0) {
                    Ordering::Less => {
                        left += 1;
                    }
                    Ordering::Greater => {
                        right -= 1;
                    }
                    Ordering::Equal => {
                        ret.push(vec![nums[i], nums[left], nums[right]]);
                        while left < right && nums[left] == nums[left + 1] {
                            left += 1;
                        }
                        while left < right && nums[right] == nums[right - 1] {
                            right -= 1;
                        }
                        left += 1;
                        right -= 1;
                    }
                }
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::three_sum(vec![-1, 0, 1, 2, -1, 4]),
        vec![vec![-1, -1, 2], vec![-1, 0, 1],]
    );
}
