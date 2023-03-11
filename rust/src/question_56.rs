struct Solution;

use std::cmp::max;

impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.sort_by(|a, b| a[0].cmp(&b[0]));

        let mut ret = Vec::new();

        let mut left = intervals[0][0];
        let mut right = intervals[0][1];
        for i in 1..intervals.len() {
            let curr_left = intervals[i][0];
            let curr_right = intervals[i][1];

            if curr_left > right {
                ret.push(vec![left, right]);
                left = curr_left;
                right = curr_right;
                continue;
            }

            right = max(right, curr_right);
        }
        ret.push(vec![left, right]);

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::merge(vec![vec![1, 3], vec![2, 6], vec![8, 10], vec![15, 18]]),
        vec![vec![1, 6], vec![8, 10], vec![15, 18]]
    );
    assert_eq!(
        Solution::merge(vec![vec![1, 4], vec![4, 5]]),
        vec![vec![1, 5]]
    );
    assert_eq!(
        Solution::merge(vec![
            vec![2, 3],
            vec![4, 5],
            vec![6, 7],
            vec![8, 9],
            vec![1, 10]
        ]),
        vec![vec![1, 10]]
    )
}
