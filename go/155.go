type MinStack struct {
    stack0 []int
    stack1 []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    stack := MinStack{stack0: []int{}, stack1: []int{}}
    return stack
}


func (this *MinStack) Push(x int)  {
    this.stack0 = append(this.stack0, x)
    if len(this.stack1) == 0 || this.GetMin() >= x {
        this.stack1 = append(this.stack1, x)
    }
}


func (this *MinStack) Pop()  {
    if this.GetMin() == this.Top() {
        this.stack1 = this.stack1[:len(this.stack1)-1]
    }
    this.stack0 = this.stack0[:len(this.stack0)-1]
}


func (this *MinStack) Top() int {
    return this.stack0[len(this.stack0)-1]
    //return len(this.stack0)
}


func (this *MinStack) GetMin() int {
    return this.stack1[len(this.stack1)-1]
    //return len(this.stack1)
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
