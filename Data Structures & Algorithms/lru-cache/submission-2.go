type Node struct {
	Key int
	Val int
	Next *Node
	Prev *Node
}

type LRUCache struct {	
    Capacity int
	Cache map[int]*Node
	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
    
	head.Next = tail
	tail.Prev = head

    return LRUCache{
		Capacity: capacity, 
		Cache: map[int]*Node{},
		Head: head,
		Tail: tail,
	}
}

func (lru *LRUCache) cacheFull() bool {
	if len(lru.Cache) == lru.Capacity{
		return true
	}
	return false
}

func (lru *LRUCache) remove(node *Node){
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) addFront(node *Node){
	node.Prev = lru.Head
	node.Next = lru.Head.Next
	lru.Head.Next.Prev = node // old first node points back to node
	lru.Head.Next = node      // must come last
}

func (lru *LRUCache) Get(key int) int {
  v, ok :=  lru.Cache[key]
  if !ok {
	return -1
  }

  lru.remove(v)
  lru.addFront(v)
  return v.Val
}

func (lru *LRUCache) Put(key int, value int) {
   // update the value if the key exist
   // add the key, value pair if the key does not exist
   // if the insertion is not possible due to the cache being full, remove the leaset used key.    
   
   _, ok := lru.Cache[key]
   if !ok {
	// insert a new node
	newEntry := &Node{
		Val: value,
		Key: key, // key is needed later to trace back to the map entry and remove the item there as well.
	}
	lru.Cache[key] = newEntry
	lru.addFront(newEntry)
    
	if len(lru.Cache) == lru.Capacity + 1 {
		// remove the least used
		node := lru.Tail.Prev
		lru.remove(node)
		delete(lru.Cache, node.Key)

	}

   }else{
     node := lru.Cache[key]	
	 node.Val = value

	 if lru.Head.Next == node {return}
	 
	 lru.remove(node)
	 lru.addFront(node)

   }
   

}
