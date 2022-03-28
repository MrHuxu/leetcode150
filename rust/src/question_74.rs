struct Solution;

use std::cmp::Ordering;

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let mut row_idx = 0;

        let mut left = 0;
        let mut right = matrix.len() - 1;
        while left <= right {
            let mid = (left + right) / 2;

            if target < matrix[mid][0] {
                if mid == 0 {
                    return false;
                }

                right = mid - 1;
                continue;
            }

            if mid < matrix.len() - 1 && target >= matrix[mid + 1][0] {
                left = mid + 1;
                continue;
            }

            row_idx = mid;
            break;
        }

        if !(row_idx < matrix.len()
            && target >= matrix[row_idx][0]
            && target <= matrix[row_idx][matrix[0].len() - 1])
        {
            return false;
        }

        left = 0;
        right = matrix[0].len() - 1;
        while left <= right {
            let mid = (left + right) / 2;

            match matrix[row_idx][mid].cmp(&target) {
                Ordering::Equal => return true,
                Ordering::Less => left += 1,
                Ordering::Greater => right -= 1,
            }
        }

        false
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::search_matrix(
            vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]],
            3
        ),
        true
    );

    assert_eq!(
        Solution::search_matrix(
            vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 60]],
            13
        ),
        false
    );
    assert_eq!(Solution::search_matrix(vec![vec![1]], 0), false);
    assert_eq!(Solution::search_matrix(vec![vec![1], vec![3]], 3), true);
}
