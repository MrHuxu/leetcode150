struct Solution;

use std::cmp::min;

impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ret = Vec::new();
        for i in 0..(min(matrix.len(), matrix[0].len()) + 1) / 2 {
            let mut tmp1 = Vec::new();
            let mut tmp2 = Vec::new();
            for j in i..matrix[0].len() - i {
                tmp1.push(matrix[i][j]);

                if matrix.len() - 1 - i != i {
                    tmp2.push(matrix[matrix.len() - 1 - i][matrix[0].len() - 1 - j]);
                }
            }

            for j in i + 1..matrix.len() - 1 - i {
                tmp1.push(matrix[j][matrix[0].len() - 1 - i]);

                if matrix[0].len() - 1 - i != i {
                    tmp2.push(matrix[matrix.len() - 1 - j][i]);
                }
            }

            ret = vec![ret, tmp1, tmp2].concat();
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::spiral_order(vec![vec![3], vec![2]]), vec![3, 2]);
    assert_eq!(
        Solution::spiral_order(vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]]),
        vec![1, 2, 3, 6, 9, 8, 7, 4, 5]
    );
    assert_eq!(
        Solution::spiral_order(vec![
            vec![1, 2, 3, 4],
            vec![5, 6, 7, 8],
            vec![9, 10, 11, 12]
        ]),
        vec![1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]
    );
}
