impl Solution {
    pub fn solve(board: &mut Vec<Vec<char>>) {
        let rows = board.len();
    let cols = board[0].len();

    for r in 0..rows {
        if board[r][0] == 'O' {
            Self::dfs_v2(board, r, 0);
        }
        if board[r][cols - 1] == 'O' {
             Self::dfs_v2(board, r, cols - 1);
        }
    }
    for c in 0..cols {
        if board[0][c] == 'O' {
             Self::dfs_v2(board, 0, c);
        }
        if board[rows - 1][c] == 'O' {
             Self::dfs_v2(board, rows - 1, c);
        }
    }

    for r in 0..rows {
        for c in 0..cols {
            if board[r][c] == 'C' {
                board[r][c] = 'O';
            } else {
                board[r][c] = 'X';
            }
        }
    }
        
    }
    pub fn dfs_v2(board: &mut Vec<Vec<char>>, r: usize, c: usize) {
    // if border_with_o(board, r, c) {
    //     return;
    // }
    let mut stack: Vec<(usize, usize)> = Vec::new();
    stack.push((r, c));
    board[r][c] = 'C';
    while stack.len() > 0 {
        let (row, col) = stack.pop().unwrap();
        let directions = [
            (row + 1, col),
            (row.saturating_sub(1), col),
            (row, col + 1),
            (row, col.saturating_sub(1)),
        ];

        for (r, c) in directions {
            if r < board.len() && c < board[0].len() && board[r][c] == 'O' {
                board[r][c] = 'C';
                stack.push((r, c));
            }
        }
    }
}
}
