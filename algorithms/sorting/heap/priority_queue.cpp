#include "priority_queue.h"

void PriorityQueue::insert(int val) {
    h.insert(val);
}

int PriorityQueue::peek() const {
    return h.get_min();
}

int PriorityQueue::extract_min() {
   return h.extract_min();
}

void PriorityQueue::decrease_key(std::vector<int> &v, int idx, int new_key) {
    h.decrease_key(idx, new_key);
}
