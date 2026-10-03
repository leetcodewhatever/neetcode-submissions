func isValidSudoku(board [][]byte) bool {

// [["1","2",".",".","3",".",".",".","."],
//  ["4","",".","5",".",".",".",".","."],
//  ["3","9","8",".",".",".",".",".","3"],
//  ["5",".",".",".","6",".",".",".","4"],
//  [".",".",".","8",".","3",".",".","5"],
//  ["7",".",".",".","2",".",".",".","6"],
//  [".",".",".",".",".",".","2",".","."],
//  [".",".",".","4","1","9",".",".","8"],
//  [".",".",".",".","8",".",".","7","9"]]


if len(board) != 9 || len(board[0]) != 9 {
	return false
}

rowSet := newHashSet[Key]()
colSet := newHashSet[Key]()
boxSet := newHashSet[Key]() 

for r := range(9){
	for c := range(9){

		if board[r][c] == '.' {
			continue
		}

		v := board[r][c]

		rk := Key{
			Row: r,
			Col: c,
			Val: v,
		}

		ck := Key{
			Row: r,
			Col: c,
			Val: v,
		}

		bk := Key{
			Row: r/3,
			Col: c/3,
			Val: v,
		}

		if rowSet.contains(rk) || colSet.contains(ck) || boxSet.contains(bk) {
			return false
		}

		rowSet.insert(rk)
		colSet.insert(ck)
		boxSet.insert(bk)
	}
}

return true
}

/*
the first observation are the following.
- for every row r, all its c pairs must be unique (1-9)
- for every column c, all it's r pairs must be unique (1-9)
- for the sub-boxes
	- on top of my head, I would scan the suboxes one by one, start with (0,0) , (3,3) and 3 along the way.


a brute force solution would be:
- fix the row every time and loop thorugh the columns.
	for r in rows
		for c in cols
- then fix  the column and loop through every row.
	for c in cols
		for r in rows
- then loop through the sub-boxes


thinkig about a better approach.
 if both rows and columns conditions are met would that means the sub boxes would be trivially correct?
		checking ... | not really..

final approach:
- we do actually need to fix the row , and colums so we can scan all of them, which is o(n^2)
- for the boxes.
	- We can change the units to boxes, meaning we can divide by 3. This would gives units in term of boxes.
	- So dividing both indexes by 3 would result in the location of that index at which box it is.
*/ 

type HashSet[T comparable] struct{
	items map[T]struct{}
}

type Key struct{
	Row, Col int
	Val byte
}

func newHashSet[T comparable]() *HashSet[T]{
	return &HashSet[T]{
		items: make(map[T]struct{}),
	}
}

func (s * HashSet[T]) insert(v T){
	s.items[v]  = struct{}{}
}

func (s * HashSet[T]) contains(v T) bool{
	_, ok := s.items[v]
	return ok 
}
