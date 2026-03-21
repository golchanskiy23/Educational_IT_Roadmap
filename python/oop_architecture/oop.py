# ООП — это парадигма программирования, в которой код организуется вокруг объектов: 
# сущностей, которые объединяют данные (атрибуты) и поведение (методы). 
# Python полностью поддерживает ООП. Четыре кита ООП: инкапсуляция, наследование, полиморфизм, абстракция.

# Класс — это шаблон (чертёж) для создания объектов. Объект — конкретный экземпляр класса. 
# Метод __init__ — конструктор: вызывается автоматически при создании объекта. 
# Первый аргумент любого метода — self — ссылка на сам объект.

class Dog:
    # общий атрибут класса
    species = "Canis lupus familiaris" 

    def __init__(self, name, breed, age):
        self.name = name
        self.breed = breed
        self.age = age

    def bark(self):
        return f"Dog with name {self.name} says: Гав"

    def info(self):
        return f"{self.name}, {self.breed}, {self.age} лет"

    def birthday(self):
        self.age += 1
        return f"Today {self.name} исполнится {self.age} лет"

class BankAccount:
    def __init__(self, owner, balance=0):
        self.owner = owner
        # "приватный" атрибут, доступен извне, если знаем манглированное имя
        self.__balance = balance
        # "защищённый" атрибут - соглашение не использовать извне - warning
        self._transactions = []

    # Свойство — доступ как атрибут, но через метод
    # аналог геттера, нужен для сокрытия реализации и удобного доступа, 
    # при этом можно реализовать вычисления и валидацию внутри
    @property
    def balance(self):
        return self.__balance

    def deposit(self, amount):
        if amount <= 0:
            raise ValueError("Сумма должна быть положительной")
        self.__balance += amount
        self._transactions.append(f"+{amount}")

    def withdraw(self, amount):
        if amount > self.__balance:
            raise ValueError("Недостаточно средств")
        self.__balance -= amount
        self._transactions.append(f"-{amount}")

# Наследование позволяет создать новый класс на основе существующего. 
class Animal:
    def __init__(self, name, sound):
        self.name = name
        self.sound = sound

    def speak(self):
        return f"{self.name}: {self.sound}!"

    def __str__(self):
        return f"Animal({self.name})"

class Cat(Animal):
    def __init__(self, name, indoor=True):
        super().__init__(name, "Мяу")
        self.indoor = indoor

    def speak(self):  # переопределение метода
        base = super().speak()
        return base + " (загадочно)"

# Полиморфизм — способность объектов разных классов отвечать на одни и те же вызовы по-разному.
# В питоне реализуется через утиную типизацию
class Shape:
    def area(self):
        raise NotImplementedError("Подкласс должен реализовать area()")

    def describe(self):
        return f"{self.__class__.__name__}: площадь = {self.area():.2f}"

class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        import math
        return math.pi * self.radius ** 2

class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height

class Triangle(Shape):
    def __init__(self, base, height):
        self.base = base
        self.height = height

    def area(self):
        return 0.5 * self.base * self.height

# Перегрузка операторов позволяет пользовательским типам данных вести себя как встроенные
class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    # Представление для print() и str()
    def __str__(self):
        return f"Vector({self.x}, {self.y})"

    # Представление для отладки
    def __repr__(self):
        return f"Vector(x={self.x}, y={self.y})"

    # v1 + v2
    def __add__(self, other):
        return Vector(self.x + other.x, self.y + other.y)

    # v1 - v2
    def __sub__(self, other):
        return Vector(self.x - other.x, self.y - other.y)

    # v * scalar
    def __mul__(self, scalar):
        return Vector(self.x * scalar, self.y * scalar)

    # scalar * v (правый операнд)
    def __rmul__(self, scalar):
        return self.__mul__(scalar)

    # v1 == v2
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    # abs(v) — длина вектора
    def __abs__(self):
        return (self.x**2 + self.y**2) ** 0.5

    # -v — отрицание
    def __neg__(self):
        return Vector(-self.x, -self.y)

from functools import total_ordering
# Полный набор операторов сравнения - модуль total_ordering
@total_ordering
class Temperature:
    def __init__(self, celsius):
        self.celsius = celsius

    @property
    def fahrenheit(self):
        return self.celsius * 9/5 + 32

    def __str__(self):
        return f"{self.celsius} градусов Цельсия"

    # для total_ordering
    def __eq__(self, other):
        if not isinstance(other, Temperature):
            raise Exception("Объект не того типа данных")
        return self.celsius == other.celsius

    def __lt__(self, other):
        if not isinstance(other, Temperature):
            raise Exception("Объект не того типа данных")
        return self.celsius < other.celsius

# Декоратор принимает функцию и возвращает функцию с доп поведением
import time

def timer(func):
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        elsapsed = time.time() - start
        print(f"Функция {func.__name__} выполнялась {elsapsed} секунд")
        print(result)
    return wrapper

@timer
def slow_timer(n: int) ->int:
    return sum(range(n))

# этот импорт позволяет сохранить метаданные функции при вызове например из main
from functools import wraps
def retry(max_attempts: int=3, delay: float=1.0):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            for i in range(1, max_attempts+1):
                try:
                    func(*args, **kwargs)
                except Exception as e:
                    print(f"Неудачная попытка номер {i}")
                    if i == max_attempts:
                        raise
                    time.sleep(delay)
        return wrapper
    return decorator

@retry(max_attempts=5, delay=2.2)
def fetch(url: str):
    """Загружает данные по URL."""
    import random
    if random.random() < 0.7:
        raise ConnectionError("Нет связи")
    return "данные"

# Декораторы класса:
# @staticmethod - метод принадлежащий самому классу
# @classmethod - метод возвращающий класс

class Ring:
    PI = 3.14159265

    def __init__(self, radius):
        self._radius = radius

    # вызов метода как атрибута, без скобок
    @property
    def radius(self):
        return self._radius

    # перехват присваивания и валидация
    @radius.setter
    def radius(self, value):
        if value < 0:
            raise ValueError("Радиус не может быть отрицательным")
        self._radius = value

    @property
    def area(self):
        return self.PI * self._radius ** 2

    @property
    def circumference(self):
        return 2 * self.PI * self._radius

    @staticmethod
    def from_diameter(diameter):
        """Фабричный метод — альтернативный конструктор."""
        return Ring(diameter / 2)

    @classmethod
    def unit_circle(cls):
        """Классовый метод — создаёт круг радиуса 1."""
        return cls(1)

# Абстрактный класс - класс, объект которого нельзя создать напрямую
# определяет набор методов, которые обязан реализовать каждый класс-наследник
from abc import ABC, abstractmethod
class AbstarctShape(ABC):
    """Абстрактный базовый класс для геометрических фигур."""
    @abstractmethod
    def area(self) -> float:
        """Вычислить площадь."""
        return 2.2
    
    @abstractmethod
    def perimeter(self) -> float:
        """Вычислить периметр."""
        return 2.2

    # Конкретный метод — одинаков для всех
    def describe(self) -> str:
        return (f"{self.__class__.__name__}: площадь={self.area():.2f}, периметр={self.perimeter():.2f}")

class CircleImpl(Shape):
    def __init__(self, radius: float):
        self.radius = radius

    def area(self) -> float:
        import math
        return math.pi * self.radius ** 2

    def perimeter(self) -> float:
        import math
        return 2 * math.pi * self.radius

class RectangleImpl(Shape):
    def __init__(self, width: float, height: float):
        self.width = width
        self.height = height

    def area(self) -> float:
        return self.width * self.height

    def perimeter(self) -> float:
        return 2 * (self.width + self.height)

# Python предоставляет богатый набор готовых ABC в модуле collections.abc: 
# Iterable, Iterator, Sequence, Mapping, MutableMapping

# Метаклассы
# Метакласс — это «класс классов»: объект, который управляет созданием самих классов. 
# В Python каждый класс является объектом типа type. type — это метакласс по умолчанию.
'''
# type — это и функция для проверки типа, и метакласс
print(type(42))      # <class "int">
print(type(int))     # <class "type">
print(type(type))    # <class "type">  — type является своим метаклассом
print(type(object))  # <class "type">
'''

# Собственный метакласс
class SingletonMeta(type):
    """Метакласс, реализующий паттерн Singleton."""
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super().__call__(*args, **kwargs)
        return cls._instances[cls]

class Database(metaclass=SingletonMeta):
    def __init__(self, url):
        self.url = url
        print(f"Подключение к {url}")

# Полезный метакласс: автоматически регистрирует подклассы
from abc import ABCMeta
class PluginMeta(ABCMeta):
    plugins = {}

    def __new__(mcs, name, bases, namespace):
        cls = super().__new__(mcs, name, bases, namespace)
        if bases and not cls.__abstractmethods__:  # не регистрируем базовый класс
            mcs.plugins[name] = cls
        return cls

class Plugin(metaclass=PluginMeta):
    @abstractmethod
    def run(self) -> str:
        return ""
        

class PluginA(Plugin):
    pass
    # def run(self): return "A"

class PluginB(Plugin):
    def run(self): return "B"


def main():
    # dog = Dog("rex", "taksa", 22)
    # print(dog.bark())
    # print(dog.bark())
    #print(dog.birthday())

    # доступ через класс и через объект
    # print(Dog.species)
    # print(dog.species)

    # acc = BankAccount("Алиса", 1000)
    # acc.deposit(500)
    # print(acc.balance)  # 1500 — через @property
    # acc.__balance → AttributeError (name mangling)
    # acc._BankAccount__balance → 1500 (обойти можно, но не надо)

    # cat = Cat("Мурка")
    # print(cat.speak())  # Мурка: Мяу! (загадочно)
    # print(isinstance(cat, Animal))
    # print(issubclass(Cat, Animal))

    # Полиморфизм в действии — один код, разные объекты
    # shapes = [Circle(5), Rectangle(4, 6), Triangle(3, 8)]

    # for shape in shapes:
    #    print(shape.describe())

    # total = sum(s.area() for s in shapes)
    # print(f"Суммарная площадь: {total:.2f}")

    # v1, v2 = Vector(1,2) , Vector(3,4)
    # print(v1+v2)
    # print(v1*3)
    # print(3*v1)
    # print(abs(v2))
    # print(-v1)
    # print(v1 == Vector(1,2))

    # temps = [Temperature(36), Temperature(100)]
    # print(sorted(temps))
    # print(max(temps))
    # print(temps[1].fahrenheit)
    # print(Temperature(20) >= Temperature(15)) # True

    # print(slow_timer(10**6))

    # print(fetch("httpru://gasporn.ru"))
    # print(fetch.__name__) # без wraps быыло бы wrapper
    # print(fetch.__doc__) # """Загружает данные по URL."""

    # c = Ring(5)
    # print(c.area)          # 78.539...
    # print(c.circumference) # 31.415...
    # c.radius = 10          # через setter

    # c2 = Ring.from_diameter(8)  # staticmethod
    # c3 = Ring.unit_circle()     # classmethod

    # shapes: list[Shape] = [CircleImpl(5), RectangleImpl(4, 6)]
    # for s in shapes:
    #    print(s.describe())

    # Создание класса через type вручную (обычно не нужно)
    # type(name, bases, namespace)
    # MyClass = type("MyClass", (object,), {"x": 10, "greet": lambda self: "Привет"})
    # obj = MyClass()
    # print(obj.greet())  # Привет
    # print(obj.x)        # 10

    db1 = Database("postgresql://localhost/mydb")
    db2 = Database("postgresql://localhost/mydb")
    print(db1 is db2)  # True — один и тот же объект
    print(PluginMeta.plugins)  # {"PluginB": <class>}

if __name__ == "__main__":
    main()