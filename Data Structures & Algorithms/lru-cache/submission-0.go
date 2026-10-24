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

func (lru *LRUCache) Get(key int) int {
  v, ok :=  lru.Cache[key]
  if !ok {
	return -1
  }
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
	
    tail := lru.Head.Next
	lru.Head.Next = newEntry
	newEntry.Next = tail
	newEntry.Prev = lru.Head
	tail.Prev = newEntry

	if len(lru.Cache) == lru.Capacity + 1 {
		// remove the least used
		remove := lru.Tail.Prev
		// head - node - remove - tail

		lru.Tail.Prev = remove.Prev
		remove.Prev.Next = lru.Tail

		delete(lru.Cache, remove.Key)

	}

   }else{
     node := lru.Cache[key]	
	 node.Val = value

	 if lru.Head.Next == node {return}
	 
	 // move the node to the front
	 // head - something - whatever - node - whatever2 - tail

	 prev := node.Prev
	 next := node.Next

	 // remove the node from the current position

	 prev.Next = next
	 next.Prev = prev

	 // insert the node to the front

      headNext := lru.Head.Next
	  
	  lru.Head.Next = node
	  node.Prev = lru.Head
	  node.Next = headNext

   }
   

}
