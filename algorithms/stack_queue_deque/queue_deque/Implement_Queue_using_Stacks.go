/*
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
-int pop() Removes the element from the front of the queue and returns it.
-int peek() Returns the element at the front of the queue.
-boolean empty() Returns true if the queue is empty, false otherwise.

Notes:
You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/

type MyQueue struct {
    in  []int
    out []int
}

func Constructor() MyQueue {
    return MyQueue{
        in:  []int{},
        out: []int{},
    }
}

func (this *MyQueue) Push(x int) {
    this.in = append(this.in, x)
}

func (this *MyQueue) move() {
    if len(this.out) == 0 {
        for len(this.in) > 0 {
            n := len(this.in)
            val := this.in[n-1]
            this.in = this.in[:n-1]
            this.out = append(this.out, val)
        }
    }
}

func (this *MyQueue) Pop() int {
    this.move()
    n := len(this.out)
    val := this.out[n-1]
    this.out = this.out[:n-1]
    return val
}

func (this *MyQueue) Peek() int {
    this.move()
    return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
    return len(this.in) == 0 && len(this.out) == 0
}