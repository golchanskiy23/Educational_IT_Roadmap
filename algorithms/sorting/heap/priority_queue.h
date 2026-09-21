#ifndef PRIORITY_QUEUE_H_
#define PRIORITY_QUEUE_H_

#include "heap.h"

// шаблон класса, компаратор - ?
class PriorityQueue {
public:
  void insert(int);
  int peek() const;
  int extract_min();
  void decrease_key(std::vector<int>&, int, int);
private:
  Heap h;  
};

#endif
