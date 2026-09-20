func longestConsecutive(nums []int) int {
/*

create a map map[int][]int.
the key is the number.
the value is the ones that are greater then the key by 1
so you do search the map with key+1.
then return the value with the biggest length.
*/

// create a hash set from the number array
// start a loop on the array.
// if n - 1 do no exist, start counting it's consecutives.

numSet := newHashSet[int]()


for _, v := range nums {
	numSet.insert(v)
}

maxC := 0
for _, v := range nums {
	if !numSet.contains(v-1){
		c := countConsecutives(numSet, v)
		if c > maxC {
			maxC = c
		}
	}
}

return maxC
}

func countConsecutives(numSet * HashSet[int], v int) int {
     count := 1
	 index := v
	 for numSet.contains(index + 1) {
		count ++
		index = index + 1
	 }
	 return count
}


type HashSet[T comparable] struct{
	items map[T]struct{}
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
