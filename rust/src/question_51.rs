struct Solution;

impl Solution {
    pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
        Self::dfs(0, &mut vec![vec!['.' as u8; n as usize]; n as usize])
    }

    fn dfs(i: usize, board: &mut Vec<Vec<u8>>) -> Vec<Vec<String>> {
        let len = board.len();

        if i == len {
            let mut ret = Vec::new();
            for row in board {
                ret.push(String::from_utf8_lossy(&row.clone()).to_string());
            }
            return vec![ret];
        }

        let mut ret = Vec::new();

        for j in 0..len {
            if !Self::validate(i, j, board) {
                continue;
            }

            board[i][j] = 'Q' as u8;
            for item in Self::dfs(i + 1, board) {
                ret.push(item);
            }
            board[i][j] = '.' as u8;
        }

        ret
    }

    fn validate(i: usize, j: usize, board: &Vec<Vec<u8>>) -> bool {
        for idx in 0..i {
            if board[idx][j] == 'Q' as u8 {
                return false;
            }
        }

        let mut step = 1;
        while i >= step && j >= step {
            if board[i - step][j - step] == 'Q' as u8 {
                return false;
            }

            step += 1;
        }

        step = 1;
        while i >= step && (j + step) < board.len() {
            if board[i - step][j + step] == 'Q' as u8 {
                return false;
            }

            step += 1;
        }

        true
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::solve_n_queens(4),
        vec![
            vec![
                String::from(".Q.."),
                String::from("...Q"),
                String::from("Q..."),
                String::from("..Q.")
            ],
            vec![
                String::from("..Q."),
                String::from("Q..."),
                String::from("...Q"),
                String::from(".Q..")
            ]
        ]
    );
}
