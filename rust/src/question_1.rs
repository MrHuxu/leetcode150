pub struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut ret = vec![0, 0];

        let mut idx: HashMap<i32, usize> = HashMap::new();
        for (i, num) in nums.iter().enumerate() {
            if idx.contains_key(&(target - num)) {
                ret = vec![*idx.get(&(target - num)).unwrap() as i32, i as i32];
                break;
            }

            idx.insert(*num, i);
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![0, 1]);
}
