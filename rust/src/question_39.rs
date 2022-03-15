struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut candidates = candidates;
        candidates.sort();
        Self::helper(0, &candidates, target)
    }

    fn helper(idx: usize, candidates: &Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut ret = Vec::new();
        for i in idx..candidates.len() {
            match candidates[i].cmp(&target) {
                Ordering::Less => {
                    for combination in Self::helper(i, candidates, target - candidates[i]) {
                        ret.push(vec![vec![candidates[i]], combination].concat())
                    }
                }
                Ordering::Equal => {
                    ret.push(vec![candidates[i]]);
                }
                Ordering::Greater => {
                    break;
                }
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::combination_sum(vec![2, 3, 6, 7], 7),
        vec![vec![2, 2, 3], vec![7]]
    );
    assert_eq!(
        Solution::combination_sum(vec![2, 3, 5], 8),
        vec![vec![2, 2, 2, 2], vec![2, 3, 3], vec![3, 5]]
    );
}
