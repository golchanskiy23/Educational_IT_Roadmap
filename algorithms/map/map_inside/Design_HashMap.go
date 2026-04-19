/*
Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
- void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
- int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
- void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
*/

type Pair struct {
    key   int
    value int
}

type MyHashMap struct {
    buckets [][]Pair
    size    int
}

func Constructor() MyHashMap {
    size := 2069
    return MyHashMap{
        buckets: make([][]Pair, size),
        size:    size,
    }
}

func (this *MyHashMap) hash(key int) int {
    return key % this.size
}

func (this *MyHashMap) Put(key int, value int) {
    index := this.hash(key)
    bucket := this.buckets[index]

    for i, pair := range bucket {
        if pair.key == key {
            this.buckets[index][i].value = value
            return
        }
    }

    this.buckets[index] = append(this.buckets[index], Pair{key, value})
}

func (this *MyHashMap) Get(key int) int {
    index := this.hash(key)
    bucket := this.buckets[index]

    for _, pair := range bucket {
        if pair.key == key {
            return pair.value
        }
    }

    return -1
}

func (this *MyHashMap) Remove(key int) {
    index := this.hash(key)
    bucket := this.buckets[index]

    for i, pair := range bucket {
        if pair.key == key {
            this.buckets[index] = append(bucket[:i], bucket[i+1:]...)
            return
        }
    }
}