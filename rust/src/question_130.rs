struct Solution;

impl Solution {
    pub fn solve(board: &mut Vec<Vec<char>>) {
        let mut visited = vec![vec![false; board[0].len()]; board.len()];
        for j in 0..board[0].len() {
            Self::dfs(board, 0, j, &mut visited);
            if board.len() - 1 != 0 {
                Self::dfs(board, board.len() - 1, board[0].len() - 1 - j, &mut visited);
            }
        }
        for i in 1..board.len() - 1 {
            Self::dfs(board, i, board[0].len() - 1, &mut visited);
            if board[0].len() - 1 != 0 {
                Self::dfs(board, board.len() - 1 - i, 0, &mut visited);
            }
        }

        for i in 0..board.len() {
            for j in 0..board[i].len() {
                if board[i][j] == 'O' && !visited[i][j] {
                    board[i][j] = 'X';
                }
            }
        }
    }

    fn dfs(board: &mut Vec<Vec<char>>, i: usize, j: usize, visited: &mut Vec<Vec<bool>>) {
        if board[i][j] == 'X' {
            return;
        }
        visited[i][j] = true;

        if i > 0 && !visited[i - 1][j] {
            Self::dfs(board, i - 1, j, visited);
        }
        if i < board.len() - 1 && !visited[i + 1][j] {
            Self::dfs(board, i + 1, j, visited);
        }
        if j > 0 && !visited[i][j - 1] {
            Self::dfs(board, i, j - 1, visited)
        }
        if j < board[i].len() - 1 && !visited[i][j + 1] {
            Self::dfs(board, i, j + 1, visited)
        }
    }
}

#[test]
fn test() {
    let mut board = vec![
        vec!['X', 'X', 'X', 'X'],
        vec!['X', 'O', 'O', 'X'],
        vec!['X', 'X', 'O', 'X'],
        vec!['X', 'O', 'X', 'X'],
    ];
    Solution::solve(&mut board);
    assert_eq!(
        board,
        vec![
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'X', 'X', 'X'],
            vec!['X', 'O', 'X', 'X'],
        ]
    );

    board = vec![vec!['O']];
    Solution::solve(&mut board);
    assert_eq!(board, vec![vec!['O']]);
}
