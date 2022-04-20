struct Solution;

use std::cmp::min;
impl Solution {
    pub fn minimum_total(triangle: Vec<Vec<i32>>) -> i32 {
        let mut triangle = triangle;

        for i in 1..triangle.len() {
            for j in 0..triangle[i].len() {
                triangle[i][j] += match (j == 0, j == triangle[i].len() - 1) {
                    (true, false) => triangle[i - 1][j],
                    (false, true) => triangle[i - 1][triangle[i - 1].len() - 1],
                    (_, _) => min(triangle[i - 1][j - 1], triangle[i - 1][j]),
                };
            }
        }

        *triangle
            .last()
            .unwrap()
            .into_iter()
            .min_by(|a, b| a.cmp(&b))
            .unwrap()
    }
}

#[test]
fn name() {
    assert_eq!(
        Solution::minimum_total(vec![vec![2], vec![3, 4], vec![6, 5, 7], vec![4, 1, 8, 3]]),
        11
    );
}
