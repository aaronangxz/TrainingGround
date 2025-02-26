package main

type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.mins) == 0 || val <= this.GetMin() {
		this.mins = append(this.mins, val)
	}
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
