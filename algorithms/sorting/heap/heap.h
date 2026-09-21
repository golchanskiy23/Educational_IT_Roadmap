#ifndef HEAP_H_
#define HEAP_H_

#include <vector>

// шаблонизировать, навесить компаратор(для возрастающей/неубывающей кучи)
class Heap {
// перераспределить приватные функции, чтобы priority queue могла их использовать
public:
  // правило пяти - ?
  Heap();

  // убрать в отдельную функцию вне класса
  std::vector<int> heap_sort(std::vector<int>& v);
private:
  void build_min_heap(std::vector<int>& v);
  int parent(int i);
  int left(int i);
  int right(int i);
  void min_heapify(std::vector<int>& v, int idx);
  // объект-массив
  std::vector<int> v;
  // число элементов в массиве и размер кучи
  // в чём разница использования size_t и int
  int arr_size, heap_size;
};

#endif // HEAP_H_
