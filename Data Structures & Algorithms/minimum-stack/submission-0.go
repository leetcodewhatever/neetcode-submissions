type MinStack struct {
	items  []int
	helper []int
}

/*
	if the stack is empty: stack.push(2)
	val = 0
    if val < stack[-1]{
	 temp = stack.top()
	 stack.pop()
	 stack.push(val)
	 satack.push(temp)
	}


*/

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)

	lastIndex := len(this.helper) - 1

	if len(this.helper) == 0 {
		this.helper = append(this.helper, val)
	} else {
		if val <= this.helper[lastIndex] {
			this.helper = append(this.helper, val)
		}
	}
}

func (this *MinStack) Pop() {
	lastItem := len(this.items) - 1
    lastItemh := len(this.helper) - 1

	if len(this.items) == 0 ||  len(this.helper) == 0{
		return
	}

	if this.items[lastItem] == this.helper[lastItemh] {
		this.helper = this.helper[:lastItemh]
	}

	this.items = this.items[:lastItem]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	lastItem := len(this.helper) - 1
    // if len(this.helper) == 0 {
    //     return -1
    // }
	return this.helper[lastItem]
}



/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */