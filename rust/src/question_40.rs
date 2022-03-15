struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn combination_sum2(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut candidates = candidates;
        candidates.sort();
        Self::helper(0, &candidates, target)
    }

    fn helper(idx: usize, candidates: &Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut ret = Vec::new();
        for i in idx..candidates.len() {
            if i != idx && i > 0 && candidates[i] == candidates[i - 1] {
                continue;
            }

            match candidates[i].cmp(&target) {
                Ordering::Less => {
                    for combination in Self::helper(i + 1, candidates, target - candidates[i]) {
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
        Solution::combination_sum2(vec![10, 1, 2, 7, 6, 1, 5], 8),
        vec![vec![1, 1, 6], vec![1, 2, 5], vec![1, 7], vec![2, 6]]
    );
    assert_eq!(
        Solution::combination_sum2(vec![2, 5, 2, 1, 2], 5),
        vec![vec![1, 2, 2], vec![5]]
    );
    assert_eq!(
        Solution::combination_sum2(vec![3, 1, 3, 5, 1, 1], 8),
        vec![vec![1, 1, 1, 5], vec![1, 1, 3, 3], vec![3, 5]]
    );
}
