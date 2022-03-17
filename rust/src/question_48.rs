struct Solution;

impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let len = matrix.len();

        if len <= 1 {
            return;
        }

        for i in 0..(len + 1) / 2 {
            for j in i..=len - 1 - i - 1 {
                let tmp = matrix[i][j];
                matrix[i][j] = matrix[len - 1 - j][i];
                matrix[len - 1 - j][i] = matrix[len - 1 - i][len - 1 - j];
                matrix[len - 1 - i][len - 1 - j] = matrix[j][len - 1 - i];
                matrix[j][len - 1 - i] = tmp;
            }
        }
    }
}

#[test]
fn test() {
    let mut matrix = vec![vec![1]];
    Solution::rotate(&mut matrix);
    assert_eq!(matrix, vec![vec![1]]);

    matrix = vec![vec![1, 2, 3], vec![4, 5, 6], vec![7, 8, 9]];
    Solution::rotate(&mut matrix);
    assert_eq!(matrix, vec![vec![7, 4, 1], vec![8, 5, 2], vec![9, 6, 3]]);

    matrix = vec![
        vec![5, 1, 9, 11],
        vec![2, 4, 8, 10],
        vec![13, 3, 6, 7],
        vec![15, 14, 12, 16],
    ];
    Solution::rotate(&mut matrix);
    assert_eq!(
        matrix,
        vec![
            vec![15, 13, 2, 5],
            vec![14, 3, 4, 1],
            vec![12, 6, 8, 9],
            vec![16, 7, 10, 11]
        ]
    );
}
