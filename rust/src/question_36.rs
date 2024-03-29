struct Solution;

impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut row = vec![vec![false; 9]; 9];
        let mut col = vec![vec![false; 9]; 9];
        let mut rect = vec![vec![false; 9]; 9];
        for i in 0..9 {
            for j in 0..9 {
                if board[i][j] == '.' {
                    continue;
                }

                let ch_idx = (board[i][j] as u8 - '0' as u8 - 1) as usize;
                let rect_idx = i / 3 * 3 + j / 3;

                if row[i][ch_idx] || col[j][ch_idx] || rect[rect_idx][ch_idx] {
                    return false;
                }

                row[i][ch_idx] = true;
                col[j][ch_idx] = true;
                rect[rect_idx][ch_idx] = true;
            }
        }

        true
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::is_valid_sudoku(vec![
            vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
            vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
            vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
            vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
            vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
            vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
            vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
            vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
            vec!['.', '.', '.', '.', '8', '.', '.', '7', '9']
        ]),
        true
    );
}
