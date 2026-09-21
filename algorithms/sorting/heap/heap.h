#ifndef HEAP_H_
#define HEAP_H_

#include <vector>

// шаблонизировать, навесить компаратор(для возрастающей/неубывающей кучи)
class Heap {
// перераспределить приватные функции, чтобы priority queue могла их использовать
public:
  // правило пяти - ?
  Heap();
  void build(std::vector<int> &v);
  bool empty() const;
  int extract_min();
private:
  int parent(int i);
  int left(int i);
  int right(int i);
  void min_heapify(std::vector<int>& v, int idx);
  // объект-массив
  std::vector<int> v;
  // размер кучи
  // в чём разница использования size_t и int?
  int heap_size;
};

#endif // HEAP_H_
