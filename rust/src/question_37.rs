struct Solution;

impl Solution {
    pub fn solve_sudoku(board: &mut Vec<Vec<char>>) {
        Self::dfs(0, 0, board);
        return;
    }

    fn dfs(i: usize, j: usize, board: &mut Vec<Vec<char>>) -> bool {
        if i == 9 {
            return true;
        }

        if j == 9 {
            return Self::dfs(i + 1, 0, board);
        }

        if board[i][j] != '.' {
            return Self::dfs(i, j + 1, board);
        }

        for num in 0..9 {
            let ch = (num + '0' as u8 + 1) as char;
            if !Self::validate(i, j, board, ch) {
                continue;
            }

            board[i][j] = ch;
            if Self::dfs(i, j + 1, board) {
                return true;
            }
            board[i][j] = '.';
        }

        false
    }

    fn validate(i: usize, j: usize, board: &Vec<Vec<char>>, ch: char) -> bool {
        for c in board[i].iter() {
            if *c == ch {
                return false;
            }
        }

        for row in board {
            if row[j] == ch {
                return false;
            }
        }

        for m in i / 3 * 3..(i / 3 + 1) * 3 {
            for n in j / 3 * 3..(j / 3 + 1) * 3 {
                if board[m][n] == ch {
                    return false;
                }
            }
        }

        true
    }
}

#[test]
fn test() {
    let mut board = vec![
        vec!['5', '3', '.', '.', '7', '.', '.', '.', '.'],
        vec!['6', '.', '.', '1', '9', '5', '.', '.', '.'],
        vec!['.', '9', '8', '.', '.', '.', '.', '6', '.'],
        vec!['8', '.', '.', '.', '6', '.', '.', '.', '3'],
        vec!['4', '.', '.', '8', '.', '3', '.', '.', '1'],
        vec!['7', '.', '.', '.', '2', '.', '.', '.', '6'],
        vec!['.', '6', '.', '.', '.', '.', '2', '8', '.'],
        vec!['.', '.', '.', '4', '1', '9', '.', '.', '5'],
        vec!['.', '.', '.', '.', '8', '.', '.', '7', '9'],
    ];
    Solution::solve_sudoku(&mut board);
    assert_eq!(
        board,
        vec![
            vec!['5', '3', '4', '6', '7', '8', '9', '1', '2'],
            vec!['6', '7', '2', '1', '9', '5', '3', '4', '8'],
            vec!['1', '9', '8', '3', '4', '2', '5', '6', '7'],
            vec!['8', '5', '9', '7', '6', '1', '4', '2', '3'],
            vec!['4', '2', '6', '8', '5', '3', '7', '9', '1'],
            vec!['7', '1', '3', '9', '2', '4', '8', '5', '6'],
            vec!['9', '6', '1', '5', '3', '7', '2', '8', '4'],
            vec!['2', '8', '7', '4', '1', '9', '6', '3', '5'],
            vec!['3', '4', '5', '2', '8', '6', '1', '7', '9']
        ]
    );
}
