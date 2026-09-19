impl Solution {
    pub fn islands_and_treasure(grid: &mut Vec<Vec<i32>>) {
let rows = grid.len();
    let cols = grid[0].len();
    let inf = 2147483647;
    let mut queue: VecDeque<(usize, usize)> = VecDeque::new();

    // seed queue
    for r in 0..rows {
        for c in 0..cols {
            if grid[r][c] == 0 {
                queue.push_back((r, c));
            }
        }
    }

    // start a multi source bfs
    while queue.len() > 0 {
        let (row, col) = queue.pop_front().unwrap();
        let directions = [
            (row + 1, col),
            (row - 1, col),
            (row, col + 1),
            (row, col - 1),
        ];

        for (r, c) in directions {
            if r < rows && c < cols && grid[r][c] == inf {
                grid[r][c] = grid[row][col] + 1;
                queue.push_back((r, c));
            }
        }
    }
    }
}
