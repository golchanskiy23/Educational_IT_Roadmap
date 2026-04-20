/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

-MinStack() initializes the stack object.
-void push(int val) pushes the element val onto the stack.
-void pop() removes the element on the top of the stack.
-int top() gets the top element of the stack.
-int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.
*/

type MinStack struct {
    stack []int
    size int
    min_stack []int
}


func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        min_stack: []int{},
        size: 0,
    }
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
    if len(this.min_stack) == 0{
        this.min_stack = append(this.min_stack, val)
    } else{
        new_min := min(this.GetMin(), val)
        this.min_stack = append(this.min_stack, new_min)
    }
    this.size++
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:this.size-1]
    this.min_stack = this.min_stack[:this.size-1]
    this.size--
}


func (this *MinStack) Top() int {
    return this.stack[this.size-1]
}


func (this *MinStack) GetMin() int {
    return this.min_stack[this.size-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */