#include "heap.h"

Heap::Heap() {}

inline int Heap::parent(int i) {
    // выкидывать ошибку?
    if (i < 0 || i >= heap_size)
      return -1;
    return i >> 1;
}

inline int Heap::left(int i) {
  if (i < 0 || i >= heap_size) return -1;
  return i << 1;
}

inline int Heap::right(int i) {
  if (i < 0 || i >= heap_size)
    return -1;
  return (i << 1)+1;
}

void Heap::min_heapify(std::vector<int>& v, int idx) {
  int l = left(idx), r = right(idx);
  int smallest = idx;
  if (l < heap_size && l != -1 && v[l] < v[idx])
    smallest= l;
  if (r < heap_size && r != -1 && v[r] < v[smallest])
    smallest = r;
  if (smallest != idx) {
      std::swap(v[idx], v[smallest]);
      min_heapify(v, smallest);
  }
}

void Heap::build_min_heap(std::vector<int>& v) {
  heap_size = arr_size = v.size();
  for (int i = v.size() / 2 - 1; i >= 0; i--) {
      min_heapify(v, i);
  }
}

std::vector<int> Heap::heap_sort(std::vector<int> &v) {
  // обработка пустого массива !!!
  build_min_heap(v);
  std::vector<int> sorted_arr;
    for (int i = v.size() - 1; i >= 1; i--) {
        sorted_arr.push_back(v[0]);
        std::swap(v[0], v[i]);
        heap_size--;
        min_heapify(v, 0);
    }
    sorted_arr.push_back(v[0]);
    return sorted_arr;
}
