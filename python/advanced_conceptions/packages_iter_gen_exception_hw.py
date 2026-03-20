import math
import operator
import os

# Напишите генераторную функцию primes(limit), которая лениво выдаёт простые числа до limit (не включая). 
# Простое число — делится только на 1 и на себя. 
# Проверку простоты оформите отдельной вспомогательной функцией is_prime(n). 

def is_prime(n: int) ->bool:
    for i in range(2, math.isqrt(n)+1):
        if n % i == 0 and n != i:
            return False
    return True

# генераторная функция
def primes(limit: int):
    for i in range(2,limit+1):
        if is_prime(i):
            yield i

from itertools import islice
def infinite_primes():
    n = 2
    while True:
        if is_prime(n):
            yield n
        n += 1

# Напишите функцию safe_calc(expression), которая принимает строку вида "10 / 2" (два числа и операция) 
# и возвращает результат. Функция должна корректно обрабатывать все возможные ошибки через try/except, 
# а не через условия.

class CalcError(ValueError):
    pass

ops = {"+": operator.add, "-": operator.sub, "*": operator.mul, "/": operator.truediv}

def safe_calc(expression: str) ->int:
    parts = expression.split()
    if len(parts) != 3:
        raise CalcError("Неверный формат...")
    elif parts[1] not in ("+", "-", "*", "/"):
        raise CalcError("Неизвестная операция")

    try:
        a,b = int(parts[0]), int(parts[2])
    except ValueError:
        raise CalcError("Не числа")
    
    if parts[1] == "/" and b == 0:
        raise CalcError("Деление на ноль")
    
    return ops[parts[1]](a, b)

# Напишите модуль tree.py с функцией print_tree(path, indent=0), которая рекурсивно выводит 
# содержимое директории в виде дерева (как команда tree в терминале). Дополнительно реализуйте 
# генераторную версию для сбора статистики.

'''
 Требуемый вывод для папки:
 documents/
 ├── report.txt         (4.2 KB)
 ├── data/
 │   ├── input.csv      (15.8 KB)
 │   └── output.csv     (8.1 KB)
 └── notes.txt          (0.9 KB)
'''

from pathlib import Path
from collections import defaultdict
def print_tree(path: Path, prefix: str = ""):
    entries = list(path.iterdir())

    for i, entry in enumerate(entries):
        connector = "└── " if i == len(entries)-1 else "|──" 
        if entry.is_file():
            size_kb = entry.stat().st_size / 1024
            print(f"{prefix}{connector}{entry.name} ({size_kb:.1f}) КБ")
        else:
            print(f"{prefix}{connector}{entry.name}/")
            extension = "   " if i == len(entries)-1 else "|   "
            try:
                print_tree(entry, prefix+extension)
            except PermissionError:
                print(f"{prefix}{extension}└──[доступ запрещён]")

# Генераторная функция
def all_files(path: Path):
    path = Path(path)
    try:
        for entry in path.iterdir():
            if entry.is_file():
                yield entry
            elif entry.is_dir():
                yield from all_files(entry)
    except PermissionError:
        return

def stats(path: Path):
    total_files, total_size, ext = 0,0, {}
    for file in all_files(path):
        total_files += 1
        total_size += file.stat().st_size / 1024
        if file.suffix not in ext:
            ext[file.suffix] = 0
        ext[file.suffix] += 1
    print(f"Total files: {total_files} , total_size: {total_size} Кб")
    for key, val in ext.items():
        print(f"{key} : {val}")

# Реализуйте класс Stack (стек — структура данных LIFO: последним вошёл, первым вышел). 
# Класс должен полностью поддерживать питоновские протоколы через магические методы, 
# а также уметь вычислять математические выражения в обратной польской нотации (ОПН).

class StackError(Exception):
    pass

class StackEmptyError(StackError):
    pass

class StackFullError(StackError):
    def __init__(self, limit):
        super().__init__(f"Стек переполнен, лимит: ({limit} элементов)")
        self.limit = limit

class Stack:
    # default size = 10
    def __init__(self, capacity=10):
        self._items = []
        self.capacity = capacity

    def __len__(self):
        return len(self._items)

    def __bool__(self):
        return len(self._items) > 0

    def __str__(self):
        return f"Stack({self._items}) top->{self._items[-1] if self._items else "None"}"

    def __repr__(self):
        return f"Stack(items={self._items} , limit = {self.capacity})"
    
    def __iter__(self):
        return iter(self._items)

    def __contains__(self, item):
        return item in self._items

    def __eq__(self, other):
        if not isinstance(other, Stack):
            return False
        return self._items == other._items

    def push(self, item):
        if len(self._items) >= self.capacity:
            raise StackFullError(self.capacity)
        self._items.append(item)

    def pop(self):
        if len(self._items) == 0:
            raise StackEmptyError("Стек пуст")
        item = self._items[-1]
        self._items = self._items[:-1]
        return item

    def top(self):
        if len(self._items) == 0:
            raise StackEmptyError("Попытка взять элемент пустого стека")
        return self._items[-1]

def evaluate(expression: str) ->int:
    stack = Stack()
    for token in expression.split(" "):
        try:
            if token in ops:
                b,a = stack.pop(), stack.pop()
                stack.push(ops[token](a,b))
            else:
                stack.push(int(token))
        except StackFullError as e:
            print(f"Ошибка: {e}")
        except StackEmptyError as e:
                print(f"Ошибка: {e}")

    return stack.pop()    


def main():
    # for p in primes(30):
    #    print(p, end=" ")

    # first_10 = list(islice(infinite_primes(), 10))
    # print(first_10)   # [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]

    # while True:
    #    in_ = input()
    #    try:
    #        print(safe_calc(in_))
    #    except CalcError as c:
    #        print(f"Ошибка: {c}")
    #        os._exit(1)

    # path = Path("/home/bloom/all/all_files")
    # print_tree(path)
    # print()
    # stats(path)

    # постфикс: "2 3 + 1 5 * 2 / -"
    print(evaluate("2 3 + 1 5 * 2 / -"))


if __name__ == "__main__":
    main()