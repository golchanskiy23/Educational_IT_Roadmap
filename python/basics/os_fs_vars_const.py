import os
import pathlib
from pathlib import Path

def os_func():
    # Текущая рабочая директория
    print(os.getcwd())           # /home/user/projects

    print(Path.cwd())            # /home/user/projects

    # Переменные окружения
    print(os.environ.get('HOME'))  # /home/user
    print(os.environ.get('PATH'))  # список директорий

    # Абсолютный путь
    p = Path('/home/user/documents/file.txt')

    # Относительный путь
    p = Path('documents/file.txt')  # относительно текущей папки

    # Существует ли путь?
    p.exists()        # True / False
    p.is_file()       # это файл?
    p.is_dir()        # это папка?

    # Части пути
    p = Path('/home/user/file.txt')
    print(p.parent)   # /home/user
    print(p.name)     # file.txt
    print(p.stem)     # file
    print(p.suffix)   # .txt

    # Содержимое папки
    for item in Path('.').iterdir():
        print(item)

    # Только файлы .py в папке
    for f in Path('.').glob('*.py'):
        print(f)


def vars():
    # Обычное присваивание
    name = 'Alice'
    age = 25
    height = 1.72

    # Множественное присваивание (bus assignment)
    a, b = 3, 4
    x, y, z = 1, 2, 3

    # Обмен значений без временной переменной
    a, b = b, a   # после: a=4, b=3

    # Присваивание одного значения нескольким переменным
    x = y = z = 0

    # Проверка типа переменной
    print(type(name))   # <class 'str'>
    print(type(age))    # <class 'int'>

    # Строка → число
    n = int('42')          # 42
    f = float('3.14')      # 3.14

    # Число → строка
    s = str(100)           # '100'
    s = str(3.14)          # '3.14'

    # Число → булево
    bool(0)    # False
    bool(1)    # True
    bool(-5)   # True (любое ненулевое — True)

    # Строка → булево
    bool('')       # False (пустая строка — False)
    bool('hello')  # True

    # Список ↔ множество
    lst = [1, 2, 2, 3]
    s = set(lst)     # {1, 2, 3}  — дубли убраны
    lst2 = list(s)   # [1, 2, 3]  — обратно в список

    # Частая ошибка: нельзя сложить int и str
    # '5' + 3  → TypeError!
    int('5') + 3     # правильно: 8

def consts():
    # Константы — соглашение, не ограничение
    PI = 3.14159265
    MAX_RETRIES = 3
    DATABASE_URL = 'postgresql://localhost/mydb'
    DEBUG = False

    x = 42          # x — int
    x = 'hello'     # теперь x — str
    x = [1, 2, 3]   # теперь x — list

def conditions_cycles():
    age = 20

    if age < 18:
        print('несовершеннолетний')
    elif age < 65:
        print('взрослый')
    else:
        print('пенсионер')

    # Тернарный оператор (однострочное условие)
    status = 'adult' if age >= 18 else 'minor'

    # Сложные условия
    x = 15
    if x > 0 and x % 2 == 0:
        print('положительное чётное')

    if x < 0 or x > 100:
        print('вне диапазона')

    if not (x == 0):
        print('не ноль')

    n = 10
    total = 0
    i = 1

    while i <= n:
        total += i
        i += 1

    print(total)  # 55

    # break — прервать цикл досрочно
    # continue — перейти к следующей итерации
    i = 0
    while True:
        i += 1
        if i % 2 == 0:
            continue  # пропустить чётные
        if i > 9:
            break     # остановиться
        print(i)  # 1, 3, 5, 7, 9

    # Перебор списка
    fruits = ['apple', 'banana', 'cherry']
    for fruit in fruits:
        print(fruit)

    # Диапазон чисел
    for i in range(5):        # 0, 1, 2, 3, 4
        print(i)

    for i in range(2, 10, 2): # 2, 4, 6, 8
        print(i)

    # Перебор с индексом
    for i, fruit in enumerate(fruits):
        print(i, fruit)  # 0 apple, 1 banana, ...

    # Перебор строки
    for char in 'hello':
        print(char)  # h, e, l, l, o

    # Перебор словаря
    d = {'a': 1, 'b': 2}
    for key, value in d.items():
        print(key, value)

def collections():
    # Создание
    lst = [1, 2, 3, 4, 5]
    mixed = [1, 'hello', 3.14, True]  # можно, но нежелательно
    empty = []

    # Доступ по индексу (с нуля)
    lst[0]   # 1 — первый
    lst[-1]  # 5 — последний
    lst[-2]  # 4 — предпоследний

    # Срезы (slicing)
    lst[1:3]   # [2, 3] — с 1 по 2 включительно
    lst[:3]    # [1, 2, 3] — первые 3
    lst[2:]    # [3, 4, 5] — с 3 до конца
    lst[::-1]  # [5, 4, 3, 2, 1] — разворот с созданием копии, .reverse() - без создания копии

    # Изменение
    lst[0] = 10         # замена элемента
    lst.append(6)       # добавить в конец
    lst.insert(1, 99)   # вставить 99 на позицию 1
    lst.remove(3)       # удалить первое вхождение 3
    lst.pop()           # удалить и вернуть последний
    lst.pop(0)          # удалить и вернуть первый

    # Информация
    len(lst)            # длина
    3 in lst            # содержит ли 3?
    lst.count(2)        # сколько раз встречается 2
    lst.index(4)        # индекс первого вхождения 4

    # Сортировка
    lst.sort()           # сортировка на месте
    sorted(lst)          # новый отсортированный список
    lst.sort(reverse=True)  # по убыванию

    # Обычный способ
    squares = []
    for i in range(10):
        squares.append(i ** 2)

    # List comprehension — то же самое в одну строку
    squares = [i ** 2 for i in range(10)]

    # С условием — только чётные квадраты
    even_sq = [i ** 2 for i in range(10) if i % 2 == 0]

    # Из другого списка
    words = ['hello', 'world', 'python']
    upper = [w.upper() for w in words]
    # ['HELLO', 'WORLD', 'PYTHON']

    import array
    arr = array.array('i', [1, 2, 3, 4])  # 'i' = int

    # NumPy array — для серьёзной работы с числами
    import numpy as np
    arr = np.array([1, 2, 3, 4])
    arr * 2          # [2, 4, 6, 8]  — поэлементно
    arr.sum()        # 10
    arr.mean()       # 2.5

    # Матрица 3x3 как список списков
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]

    # Доступ к элементу: строка, столбец
    matrix[0][0]  # 1 — левый верхний
    matrix[1][2]  # 6 — вторая строка, третий столбец

    # Обход матрицы
    for row in matrix:
        for elem in row:
            print(elem, end=' ')
        print()

    # Создание матрицы n×m заполненной нулями
    n, m = 3, 4
    zeros = [[0] * m for _ in range(n)]

    # ВНИМАНИЕ: не делайте так!
    wrong = [[0] * m] * n  # все строки — один объект!

    # NumPy — удобнее для вычислений
    import numpy as np
    A = np.array([[1, 2], [3, 4]])
    B = np.array([[5, 6], [7, 8]])
    A + B          # поэлементное сложение  
    A @ B          # матричное умножение
    A.T            # транспонирование

    # Создание
    person = {'name': 'Alice', 'age': 30, 'city': 'Moscow'}
    empty = {}

    # Доступ
    person['name']          # 'Alice'
    person.get('age')       # 30
    person.get('email', 'не задан')  # 'не задан' — если нет ключа

    # Изменение
    person['age'] = 31          # обновить
    person['email'] = 'a@b.com' # добавить новый ключ
    del person['city']          # удалить ключ
    person.pop('email')         # удалить и вернуть значение

    # Перебор
    for key in person:             # только ключи
        print(key)
    for key, val in person.items():  # пары
        print(key, '->', val)
    for val in person.values():    # только значения
        print(val)

    # Проверка наличия ключа
    'name' in person    # True

    # Слияние словарей (Python 3.9+)
    d1 = {'a': 1}
    d2 = {'b': 2}
    merged = d1 | d2    # {'a': 1, 'b': 2}

    # Dict comprehension
    squares = {x: x**2 for x in range(5)}
    # {0: 0, 1: 1, 2: 4, 3: 9, 4: 16}

    # Создание
    s = {1, 2, 3, 4, 5}
    s = set([1, 2, 2, 3, 3])  # {1, 2, 3} — дубли убраны
    empty_set = set()  # НЕ {} — это пустой словарь!

    # Операции
    s.add(6)      # добавить элемент
    s.remove(3)   # удалить (ошибка если нет)
    s.discard(3)  # удалить (без ошибки если нет)

    # Математические операции
    a = {1, 2, 3, 4}
    b = {3, 4, 5, 6}

    a | b    # объединение:     {1, 2, 3, 4, 5, 6}
    a & b    # пересечение:     {3, 4}
    a - b    # разность:        {1, 2}
    a ^ b    # симметр. разн.: {1, 2, 5, 6}

    # Быстрая проверка принадлежности (быстрее, чем список!)
    3 in a    # True

    # Удаление дублей из списка
    lst = [1, 2, 2, 3, 3, 3, 4]
    unique = list(set(lst))  # [1, 2, 3, 4]

    # Создание
    t = (1, 2, 3)
    t = 1, 2, 3      # скобки необязательны
    single = (42,)   # кортеж из одного элемента — нужна запятая!
    empty = ()

    # Доступ — как у списка
    t[0]     # 1
    t[-1]    # 3
    t[1:3]   # (2, 3)

    # Неизменяемость
    # t[0] = 99  → TypeError: tuple не поддерживает присваивание

    # Распаковка
    x, y, z = (10, 20, 30)
    first, *rest = (1, 2, 3, 4, 5)
    # first = 1, rest = [2, 3, 4, 5]

    # Кортеж как ключ словаря (список нельзя!)
    coords = {}
    coords[(55.7, 37.6)] = 'Moscow'

    # Именованные кортежи — кортеж с именами полей
    from collections import namedtuple
    Point = namedtuple('Point', ['x', 'y'])
    p = Point(10, 20)
    print(p.x, p.y)  # 10 20

def ds_copiing():
    # Числа и строки — «безопасны»
    # Они неизменяемы, поэтому проблемы не возникает
    a = 5
    b = a
    b = 10
    print(a)  # 5 — не изменился

    # Списки — изменяемы! Вот где ловушка
    a = [1, 2, 3]
    b = a          # b — это тот же список, не копия!
    b.append(4)
    print(a)  # [1, 2, 3, 4]  — a тоже изменился!
    print(b)  # [1, 2, 3, 4]

    # Убедиться: a и b — один объект
    print(a is b)  # True

    import copy

    original = [1, 2, 3]

    # Три равнозначных способа shallow copy для списка:
    b = original.copy()
    b = list(original)
    b = original[:]

    b.append(99)
    print(original)  # [1, 2, 3]   — не тронут
    print(b)         # [1, 2, 3, 99]

    # Но! При вложенных объектах shallow copy не спасает
    nested = [[1, 2], [3, 4]]
    copy_n = nested.copy()   # shallow

    copy_n[0].append(99)   # изменяем внутренний список
    print(nested)    # [[1, 2, 99], [3, 4]] — оригинал пострадал!
    print(copy_n)    # [[1, 2, 99], [3, 4]]

    nested = [[1, 2], [3, 4]]
    deep = copy.deepcopy(nested)

    deep[0].append(99)
    print(nested)  # [[1, 2], [3, 4]]   — оригинал цел!
    print(deep)    # [[1, 2, 99], [3, 4]]

    # То же самое работает для словарей
    d = {'a': [1, 2, 3], 'b': [4, 5, 6]}
    d_copy = copy.deepcopy(d)
    d_copy['a'].append(99)
    print(d)       # {'a': [1, 2, 3], 'b': [4, 5, 6]}
    print(d_copy)  # {'a': [1, 2, 3, 99], 'b': [4, 5, 6]}

    