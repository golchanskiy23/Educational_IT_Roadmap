#ifndef HEAP_H_
#define HEAP_H_

#include <vector>

// шаблонизировать, навесить компаратор(для возрастающей/неубывающей кучи)
class Heap {
// перераспределить приватные функции, чтобы priority queue могла их использовать
public:
  // правило пяти - ?
  Heap();
  void build(std::vector<int> &);
  bool empty() const;
  int extract_min();
  int get_min() const;
  void insert(int);
  void decrease_key(int, int);
private:
  int parent(int);
  int left(int);
  int right(int);
  void min_heapify(std::vector<int>&, int);
  // объект-массив
  std::vector<int> v;
  // размер кучи
  // в чём разница использования size_t и int?
  int heap_size;
};

#endif // HEAP_H_
