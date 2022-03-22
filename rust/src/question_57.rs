struct Solution;

use std::cmp::max;

impl Solution {
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        let mut intervals = intervals;
        intervals.push(new_interval);
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
        Solution::insert(vec![vec![1, 3], vec![6, 9]], vec![2, 5]),
        vec![vec![1, 5], vec![6, 9]]
    );

    assert_eq!(
        Solution::insert(
            vec![
                vec![1, 2],
                vec![3, 5],
                vec![6, 7],
                vec![8, 10],
                vec![12, 16]
            ],
            vec![4, 8]
        ),
        vec![vec![1, 2], vec![3, 10], vec![12, 16]]
    );
}
