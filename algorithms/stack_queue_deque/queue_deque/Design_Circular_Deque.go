/*
Design your implementation of the circular double-ended queue (deque).

Implement the MyCircularDeque class:

-MyCircularDeque(int k) Initializes the deque with a maximum size of k.
-boolean insertFront() Adds an item at the front of Deque. Returns true if the operation is successful, or false otherwise.
-boolean insertLast() Adds an item at the rear of Deque. Returns true if the operation is successful, or false otherwise.
-boolean deleteFront() Deletes an item from the front of Deque. Returns true if the operation is successful, or false otherwise.
-boolean deleteLast() Deletes an item from the rear of Deque. Returns true if the operation is successful, or false otherwise.
-int getFront() Returns the front item from the Deque. Returns -1 if the deque is empty.
-int getRear() Returns the last item from Deque. Returns -1 if the deque is empty.
-boolean isEmpty() Returns true if the deque is empty, or false otherwise.
-boolean isFull() Returns true if the deque is full, or false otherwise.
*/

type MyCircularDeque struct {
    data  []int
    front int
    rear  int
    size  int
    cap   int
}

func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        data:  make([]int, k),
        front: 0,
        rear:  0,
        size:  0,
        cap:   k,
    }
}

func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }

    this.front = (this.front - 1 + this.cap) % this.cap
    this.data[this.front] = value
    this.size++
    return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }

    this.data[this.rear] = value
    this.rear = (this.rear + 1) % this.cap
    this.size++
    return true
}

func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }

    this.front = (this.front + 1) % this.cap
    this.size--
    return true
}

func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }

    this.rear = (this.rear - 1 + this.cap) % this.cap
    this.size--
    return true
}

func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.front]
}

func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    idx := (this.rear - 1 + this.cap) % this.cap
    return this.data[idx]
}

func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
    return this.size == this.cap
}