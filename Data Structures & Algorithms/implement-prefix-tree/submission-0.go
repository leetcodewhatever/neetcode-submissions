type TrieNode struct {
	childrens map[rune]*TrieNode
	isEnd bool
}
type PrefixTree struct {
    root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &TrieNode{
			  childrens: make(map[rune]*TrieNode),
     	},
	}  
}

func (this *PrefixTree) Insert(word string) {
    curr := this.root
	for _, c := range word {
		child, exists := curr.childrens[c]
		if !exists{
           child = &TrieNode{
			  childrens: make(map[rune]*TrieNode),
		   }
		   curr.childrens[c] = child
		}
		curr = child
	}
	
	curr.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
    curr := this.root
	for _, c := range word {
		child, exists := curr.childrens[c]
		if !exists {
			return false
		}
		curr = child
	}

	return curr.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
   curr := this.root
   for _, c := range prefix {
		child, exists := curr.childrens[c]
		if !exists{
			return false
		}
		curr = child
   }

   return true
}
