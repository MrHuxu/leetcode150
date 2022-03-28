struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let mut change_rows = HashMap::new();
        let mut change_cols = HashMap::new();

        let row_cnt = matrix.len();
        let col_cnt = matrix[0].len();
        for i in 0..row_cnt {
            for j in 0..col_cnt {
                if matrix[i][j] != 0 {
                    continue;
                }
                change_rows.insert(i, true);
                change_cols.insert(j, true);
            }
        }
        for i in 0..row_cnt {
            for j in 0..col_cnt {
                if change_rows.contains_key(&i) || change_cols.contains_key(&j) {
                    matrix[i][j] = 0;
                }
            }
        }
    }
}

#[test]
fn test() {
    let mut matrix = vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]];
    Solution::set_zeroes(&mut matrix);
    assert_eq!(matrix, vec![vec![1, 0, 1], vec![0, 0, 0], vec![1, 0, 1]]);

    matrix = vec![vec![0, 1, 2, 0], vec![3, 4, 5, 2], vec![1, 3, 1, 5]];
    Solution::set_zeroes(&mut matrix);
    assert_eq!(
        matrix,
        vec![vec![0, 0, 0, 0], vec![0, 4, 5, 0], vec![0, 3, 1, 0]]
    );
}
