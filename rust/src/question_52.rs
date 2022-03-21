struct Solution;

impl Solution {
    pub fn total_n_queens(n: i32) -> i32 {
        Self::dfs(0, &mut vec![vec!['.'; n as usize]; n as usize])
    }

    fn dfs(i: usize, board: &mut Vec<Vec<char>>) -> i32 {
        let len = board.len();

        if i == len {
            return 1;
        }

        let mut ret = 0;

        for j in 0..len {
            if !Self::validate(i, j, board) {
                continue;
            }

            board[i][j] = 'Q';
            ret += Self::dfs(i + 1, board);
            board[i][j] = '.';
        }

        ret
    }

    fn validate(i: usize, j: usize, board: &Vec<Vec<char>>) -> bool {
        for idx in 0..i {
            if board[idx][j] == 'Q' {
                return false;
            }
        }

        let mut step = 1;
        while i >= step && j >= step {
            if board[i - step][j - step] == 'Q' {
                return false;
            }

            step += 1;
        }

        step = 1;
        while i >= step && (j + step) < board.len() {
            if board[i - step][j + step] == 'Q' {
                return false;
            }

            step += 1;
        }

        true
    }
}

#[test]
fn test() {
    assert_eq!(Solution::total_n_queens(4), 2);
    assert_eq!(Solution::total_n_queens(1), 1);
}
