struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        Self::dfs(&nums, &mut HashMap::new())
    }

    fn dfs(nums: &Vec<i32>, used: &mut HashMap<i32, bool>) -> Vec<Vec<i32>> {
        let mut ret = Vec::new();
        for num in nums {
            if used.contains_key(num) {
                continue;
            }

            used.insert(*num, true);
            for tmp in Self::dfs(nums, used) {
                ret.push(vec![vec![*num], tmp].concat());
            }
            used.remove(num);
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
        Solution::permute(vec![1, 2, 3]),
        vec![
            vec![1, 2, 3],
            vec![1, 3, 2],
            vec![2, 1, 3],
            vec![2, 3, 1],
            vec![3, 1, 2],
            vec![3, 2, 1]
        ]
    );
    assert_eq!(
        Solution::permute(vec![0, 1]),
        vec![vec![0, 1], vec![1, 0,],]
    );
}
