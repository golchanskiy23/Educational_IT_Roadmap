# Модуль — это любой файл с расширением .py.
import utils

print(utils.greet("msx"))
print(utils.PI)

# Структура пакета:
'''
myproject/
├── main.py
└── mypackage/
    ├── __init__.py      # обязателен (можно пустой)
    ├── geometry.py      # модуль внутри пакета
    └── utils.py
'''

# Импорт из пакета
from mypackage import geometry
from mypackage.geometry import Circle
import mypackage.my_utils as mu

# pip — инструмент командной строки для установки сторонних библиотек из Python Package Index (PyPI).

# __pycache__ — это папка, автоматически создаваемая интерпретатором Python 
# для хранения скомпилированных файлов байт-кода (.pyc) импортированных модулей. 

#  Виртуальное окружение
# Создать виртуальное окружение в папке .venv
# python -m venv .venv

# Теперь pip install работает только внутри этого окружения
# pip install requests

# Деактивировать
# deactivate

'''
# Стандартная структура проекта:
myproject/
├── .venv/           # НЕ добавлять в git!
├── .gitignore       # добавить .venv/ сюда
├── requirements.txt
└── main.py
'''

'''
# у каждого пакета изолированная папка с зависимостями
# активация окружения заменяет PATH (переменная окружения, которая содержит список папок, где система ищет исполняемые файлы (команды))
myproject_old/
└── .venv/
    └── site-packages/
        └── Django 3.2   ← своя версия

myproject_new/
└── .venv/
    └── site-packages/
        └── Django 4.2   ← своя версия
'''

# Итераторы и генераторы
# Итератор - указатель на i-ый элемент iterable коллекции
# Список — итерируемый, но не итератор
lst = [1, 2, 3]
it = iter(lst)         # получаем итератор

print(next(it))        # 1
print(next(it))        # 2
print(next(it))        # 3
# next(it)             # StopIteration — итератор исчерпан

# for работает именно так под капотом:
# it = iter(lst)
# while True:
#     try: item = next(it)
#     except StopIteration: break

# Создание своего итератора через класс
class Countdown:
    def __init__(self, start):
        self.current = start

    def __iter__(self):
        return self            # итератор — сам объект

    def __next__(self):
        if self.current <= 0:
            raise StopIteration
        self.current -= 1
        return self.current + 1

for n in Countdown(3):
    print(n)           # 3, 2, 1

# Генератор (ленивый итератор)
# Выполнение объекта-генератора до yield

def countdown(n):
    """Генератор обратного отсчёта"""
    while n > 0:
        yield n         # приостанавливаемся и отдаём значение
        n -= 1          # продолжаем при следующем next()

gen = countdown(3)
print(type(gen))        # <class 'generator'>
print(next(gen))        # 3
print(next(gen))        # 2
print(next(gen))        # 1

for n in countdown(5):
    print(n)            # 5 4 3 2 1

# Генераторное выражение — ленивое (значения по требованию)
squares_gen = (x**2 for x in range(10))

# sum() сам вызывает next() в цикле, пока генератор не кончится
total = sum(x**2 for x in range(5)) # в чём плюс такой записи?

# Рекурсия
def factorial(n):
    if n <= 1:          # базовый случай — остановка
        return 1
    return n * factorial(n - 1)   # рекурсивный шаг

print(factorial(5))     # 5 * 4 * 3 * 2 * 1 = 120

# Трассировка вызовов:
# factorial(5)
#   → 5 * factorial(4)
#       → 4 * factorial(3)
#           → 3 * factorial(2)
#               → 2 * factorial(1)
#                   → 1   (базовый случай)
#               ← 2 * 1 = 2
#           ← 3 * 2 = 6
#       ← 4 * 6 = 24
#   ← 5 * 24 = 120

# Дополнение - yield from. Позволяет генератору «перепоручить» перебор другому итерируемому объекту. 
'''
documents/
├── a.txt
├── sub/
│   ├── b.txt
│   └── c.txt
└── d.txt
'''

def all_files(folder):
    for item in Path(folder).iterdir():
        if item.is_dir():
            yield from all_files(item)
        else:
            yield item

'''
**Итерация 1** — `item = a.txt`
```
all_files("documents") смотрит на a.txt
→ это файл → yield a.txt
→ наружу уходит: a.txt
```

**Итерация 2** — `item = sub/`
```
all_files("documents") смотрит на sub/
→ это папка → yield from all_files("sub")
→ "documents" засыпает, управление переходит в all_files("sub")

    all_files("sub") смотрит на b.txt
    → это файл → yield b.txt
    → наружу уходит: b.txt
```

**Итерация 3** — всё ещё внутри `sub/`
```
all_files("sub") продолжает, смотрит на c.txt
→ это файл → yield c.txt
→ наружу уходит: c.txt

all_files("sub") закончился
→ управление возвращается в all_files("documents")
→ "documents" просыпается и идёт дальше
```

**Итерация 4** — `item = d.txt`
```
all_files("documents") смотрит на d.txt
→ это файл → yield d.txt
→ наружу уходит: d.txt

all_files("documents") закончился

a.txt → b.txt → c.txt → d.txt
'''

# рекурсия : собирает список -> печатает весь список
# генератор : получает элемент -> печатает -> забывает -> ...


# Магические методы - переопределение операций, поведение аналогично встроенным типам
class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __str__(self):           # объясняем: как печатать
        return f"({self.x}, {self.y})"

    def __add__(self, other):    # объясняем: что значит +
        return Vector(self.x + other.x, self.y + other.y)

    def __len__(self):           # объясняем: что значит len()
        return 2

    def __eq__(self, other):     # объясняем: что значит ==
        return self.x == other.x and self.y == other.y

v1 = Vector(1, 2)
v2 = Vector(3, 4)

print(v1)       # (1, 2)
v1 + v2         # Vector(4, 6)
len(v1)         # 2
v1 == v2        # False — теперь сравнивает значения


'''
Главная таблица магических методов: 
__init__ (конструктор), 
__str__ (читаемое представление), 
__repr__ (отладочное), 
__len__, 
__getitem__, 
__setitem__, 
__contains__, 
__iter__, 
__add__/__sub__/__mul__ (арифметика), 
__eq__/__lt__ (сравнение), 
__enter__/__exit__ (with).
'''

# Value Error <- Exception <- Base Exception -> Системные ошибки(sys.exit() , ctrl+C)
def safe_divide(a, b):
    try:
        result = a / b
    except ZeroDivisionError:
        print("Ошибка: деление на ноль")
        return None
    except TypeError as e:
        print(f"Ошибка типа: {e}")
        return None
    else:
        print("Деление прошло успешно")
        return result
    finally:
        print("Этот блок выполняется всегда")  # даже при return

print(safe_divide(10, 2))   # Деление успешно; Этот блок...; 5.0
print(safe_divide(10, 0))   # Ошибка: деление...; Этот блок...; None

# Пользовательские исключения
# Исключение с дополнительными данными
class InsufficientFundsError(Exception):
    def __init__(self, balance, amount):
        self.balance = balance
        self.amount = amount
        super().__init__(
            f"Недостаточно средств: баланс {balance}, запрос {amount}"
        )

class BankAccount:
    def __init__(self, balance=0):
        self.balance = balance

    def withdraw(self, amount):
        if amount < 0:
            raise NegativeAmountError(f"Сумма не может быть отрицательной: {amount}")
        if amount > self.balance:
            raise InsufficientFundsError(self.balance, amount)
        self.balance -= amount
        return amount

acc = BankAccount(100)
try:
    acc.withdraw(200)
except InsufficientFundsError as e:
    print(e)                    # Недостаточно средств: баланс 100, запрос 200
    print(f"Не хватает: {e.amount - e.balance}")  # 100
except NegativeAmountError as e:
    print(f"Некорректная сумма: {e}")


# raise - явно поднять исключение
# reraise - залоггировать и пробросить исключение дальше
def read_config(path):
    try:
        with open(path) as f:
            data = f.read()
    except FileNotFoundError:
        # raise — поднять своё исключение вместо стандартного
        raise RuntimeError(f"Конфиг не найден: {path}")
    
    try:
        return json.loads(data)
    except json.JSONDecodeError as e:
        # raise ... from — цепочка: новое исключение + причина
        raise RuntimeError("Конфиг повреждён") from e

def start_app():
    try:
        config = read_config("settings.json")
    except RuntimeError:
        print("Логируем ошибку...")
        raise   # reraise — перебрасываем то же исключение дальше
                # без аргументов — исключение не меняется

try:
    start_app()
except RuntimeError as e:
    print(f"Приложение не запустилось: {e}")