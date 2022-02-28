struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
        let mut nums = nums;
        nums.sort();

        let mut ret: Option<i32> = None;
        let mut distance = 0;

        for i in 0..nums.len() {
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }

            let mut left = i + 1;
            let mut right = nums.len() - 1;
            while left < right {
                let sum = nums[i] + nums[left] + nums[right];

                match sum.cmp(&target) {
                    Ordering::Greater => {
                        if ret.is_none() || sum - target < distance {
                            ret = Some(sum);
                            distance = sum - target;
                        }
                        right -= 1;
                    }
                    Ordering::Less => {
                        if ret.is_none() || target - sum < distance {
                            ret = Some(sum);
                            distance = target - sum;
                        }
                        left += 1;
                    }
                    Ordering::Equal => {
                        return sum;
                    }
                }
            }
        }

        ret.unwrap()
    }
}

#[test]
fn test() {
    assert_eq!(Solution::three_sum_closest(vec![-1, 2, 1, -4], 1), 2);
    assert_eq!(Solution::three_sum_closest(vec![0, 0, 0], 1), 0);
    assert_eq!(Solution::three_sum_closest(vec![-3, -2, -5, 3, -4], -1), -2);
}
