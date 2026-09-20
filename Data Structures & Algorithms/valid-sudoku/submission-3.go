func isValidSudoku(board [][]byte) bool {


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
				Val: v,
			}

			ck := Key{
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

type HashSet[T comparable] struct{
	items map[T]struct{}
}

type Key struct{
	Row int
	Col int
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
