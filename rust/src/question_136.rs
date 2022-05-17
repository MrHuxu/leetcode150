struct Solution;

impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut ret = 0;
        for i in 0..nums.len() {
            ret ^= nums[i];
        }
        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::single_number(vec![1, 1, 2]), 2);
}
