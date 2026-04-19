/*
Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

-void add(key) Inserts the value key into the HashSet.
-bool contains(key) Returns whether the value key exists in the HashSet or not.
-void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.
*/

type MyHashSet struct {
    buckets [][]int
    size    int
}

func Constructor() MyHashSet {
    size := 2069
    return MyHashSet{
        buckets: make([][]int, size),
        size:    size,
    }
}

func (this *MyHashSet) hash(key int) int {
    return key % this.size
}

func (this *MyHashSet) Add(key int) {
    index := this.hash(key)
    bucket := this.buckets[index]

    for _, k := range bucket {
        if k == key {
            return
        }
    }

    this.buckets[index] = append(bucket, key)
}

func (this *MyHashSet) Remove(key int) {
    index := this.hash(key)
    bucket := this.buckets[index]

    for i, k := range bucket {
        if k == key {
            this.buckets[index] = append(bucket[:i], bucket[i+1:]...)
            return
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    index := this.hash(key)
    bucket := this.buckets[index]

    for _, k := range bucket {
        if k == key {
            return true
        }
    }
    return false
}