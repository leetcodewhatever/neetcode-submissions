type Node struct {
	children map[rune]*Node
	isEnd bool
}
type WordDictionary struct {
   root *Node 
}

func Constructor() WordDictionary {
    return WordDictionary {
		root: &Node{
			children: make(map[rune]*Node),
		},
	} 
}

func (this *WordDictionary) AddWord(word string)  {
   curr := this.root 
   for _, c := range word {
	child, ok := curr.children[c]
	if !ok {
		child = &Node{
			children: make(map[rune]*Node),
		}

		curr.children[c] = child
	}
	curr = child
   }

   curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.root, word, 0)
}

func dfs(node *Node, word string, i int) bool{
    if i == len(word){
		return node.isEnd
	}

	c := rune(word[i])

	if c == '.' {
		for _, child := range node.children {
			if dfs(child, word, i + 1) {
				return true
			}
		}
		return false
	}

	child, ok := node.children[c]
	if !ok {
		return false
	}
	return dfs(child, word, i + 1)
}
