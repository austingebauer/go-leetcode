package min_stack_155

type stackNode struct {
	val  int
	next *stackNode
}

type MinStack struct {
	top *stackNode
	min *stackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	sn := &stackNode{
		val:  x,
		next: nil,
	}

	if this.top == nil {
		this.top = sn
	} else {
		sn.next = this.top
		this.top = sn
	}

	if this.min == nil {
		this.min = sn
	} else if sn.val < this.min.val {
		this.min = sn
	}
}

func (this *MinStack) Pop() {
	if this.top == nil {
		return
	}

	if this.top == this.min {
		if this.top.next == nil {
			this.min = nil
			this.top = nil
		} else {
			this.top = this.top.next
			this.min = this.top

			minPtr := this.min
			for minPtr != nil {
				if minPtr.val < this.min.val {
					this.min = minPtr
				}
				minPtr = minPtr.next
			}
		}
		return
	}

	this.top = this.top.next
}

func (this *MinStack) Top() int {
	return this.top.val
}

func (this *MinStack) GetMin() int {
	return this.min.val
}
