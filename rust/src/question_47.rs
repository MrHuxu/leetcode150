struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn permute_unique(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut nums = nums;
        nums.sort();
        Self::dfs(&nums, &mut HashMap::new())
    }

    fn dfs(nums: &Vec<i32>, used: &mut HashMap<usize, bool>) -> Vec<Vec<i32>> {
        let mut ret = Vec::new();
        for (i, num) in nums.iter().enumerate() {
            if used.contains_key(&i)
                || (i > 0 && nums[i] == nums[i - 1] && !used.contains_key(&(i - 1)))
            {
                continue;
            }

            used.insert(i, true);
            for tmp in Self::dfs(nums, used) {
                ret.push(vec![vec![*num], tmp].concat());
            }
            used.remove(&i);
        }

        if ret.len() == 0 {
            return vec![vec![]];
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::permute_unique(vec![1, 1, 2]),
        vec![vec![1, 1, 2], vec![1, 2, 1], vec![2, 1, 1]]
    )
}
