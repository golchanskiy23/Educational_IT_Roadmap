# Реализуйте класс Matrix, представляющий математическую матрицу. 
# Класс должен поддерживать основные операции через магические методы.
class Matrix:
    def __init__(self, data: list[list]):
        self._data = [row[:] for row in data]

    @property
    def rows(self) -> int:
        return len(self._data)

    @property
    def cols(self) -> int:
        return len(self._data[0]) if self._data else 0

    def __getitem__(self, index):
        return self._data[index]

    def __str__(self) -> str:
        lines = []
        for row in self._data:
            lines.append("[ ".join(f"{v:2}" for v in row) + "]")
        return "[" + "\n ".join(lines) + "]"

    def __repr__(self) -> str:
        return f"Matrix({self._data!r})"

    def __add__(self, other: "Matrix") -> "Matrix":
        if self.rows != other.rows or self.cols != other.cols:
            raise ValueError(f"Несовместимые размеры: {self.rows}x{self.cols} и {other.rows}x{other.cols}")
        return Matrix([
            [self._data[i][j] + other._data[i][j] for j in range(self.cols)]
            for i in range(self.rows)
        ])

    def __mul__(self, other) -> "Matrix":
        if isinstance(other, (int, float)):
            return Matrix([[v * other for v in row] for row in self._data])
        if isinstance(other, Matrix):
            if self.cols != other.rows:
                raise ValueError(f"Несовместимые размеры для умножения: {self.rows}x{self.cols} * {other.rows}x{other.cols}")
            return Matrix([
                [sum(self._data[i][k] * other._data[k][j] for k in range(self.cols)) for j in range(other.cols)]
                for i in range(self.rows)
            ])
        return NotImplemented

    def __rmul__(self, scalar) -> "Matrix":
        return self.__mul__(scalar)

    def transpose(self) -> "Matrix":
        return Matrix([
            [self._data[i][j] for i in range(self.rows)]
            for j in range(self.cols)
        ])


# Создайте иерархию классов для зоопарка, используя абстрактные классы,  наследование и полиморфизм.
from abc import ABC, abstractmethod

class Animal(ABC):
    def __init__(self, name: str, age: int):
        self.name = name
        self.age = age

    @abstractmethod
    def speak(self) -> str: ...

    @abstractmethod
    def move(self) -> str: ...

    @property
    @abstractmethod
    def diet(self) -> str: ...

    def info(self) -> str:
        return f"{self.name} ({self.__class__.__name__}), {self.age} лет, {self.diet}"


class Lion(Animal):
    @property
    def diet(self) -> str:
        return "хищник"

    def speak(self) -> str:
        return "Рычание"

    def move(self) -> str:
        return "бег"


class Dolphin(Animal):
    @property
    def diet(self) -> str:
        return "рыбоед"

    def speak(self) -> str:
        return "Свист"

    def move(self) -> str:
        return "плавание"


class Eagle(Animal):
    @property
    def diet(self) -> str:
        return "хищник"

    def speak(self) -> str:
        return "Клёкот"

    def move(self) -> str:
        return "полёт"


class Zoo:
    def __init__(self, name: str):
        self.name = name
        self._animals: list[Animal] = []

    def add(self, animal: Animal):
        self._animals.append(animal)

    def remove(self, name: str):
        self._animals = [a for a in self._animals if a.name != name]

    def by_diet(self, diet: str) -> list[Animal]:
        return [a for a in self._animals if a.diet == diet]

    def oldest(self) -> Animal:
        return max(self._animals, key=lambda a: a.age)

    def __len__(self) -> int:
        return len(self._animals)

    def __iter__(self):
        return iter(self._animals)

    def __str__(self) -> str:
        return f"Зоопарк «{self.name}» ({len(self)} животных)"

# Напишите декоратор @cache(ttl=60), который кэширует результаты функции. 
# TTL (time-to-live) — время жизни кэша в секундах. 
# После истечения TTL результат пересчитывается заново.

import time
from functools import wraps

def cache(ttl=None):
    def make_cached(func, ttl_seconds=None):
        cache, hits, misses = {}, 0, 0
        @wraps(func)
        def wrapper(*args, **kwargs):
            nonlocal hits, misses
            key = (args, frozenset(kwargs.items()))
            if key in cache:
                result, ts = cache[key]
                if ttl_seconds is None or time.time() - ts <= ttl_seconds:
                    hits += 1
                    return result
            misses += 1
            result = func(*args, **kwargs)
            cache[key] = (result, time.time())
            return result

        def cache_info() -> dict:
            return {"hits": hits, "misses": misses, "size": len(cache)}

        def cache_clear():
            nonlocal hits, misses
            cache.clear()
            hits = 0
            misses = 0

        wrapper.cache_info = cache_info
        wrapper.cache_clear = cache_clear
        return wrapper

    if callable(ttl):
        func = ttl
        return make_cached(func, ttl_seconds=None)
    return lambda func: make_cached(func, ttl_seconds=ttl)

# Реализуйте упрощённый ORM (Object-Relational Mapper): систему, которая позволяет описывать 
# модели данных как Python-классы и автоматически сохранять/загружать их в JSON-файл.
import json
import os

# __set_name__ — вызывается автоматически когда класс создаётся, запоминает имя атрибута:

# __get__ и __set__ — дескрипторные методы, перехватывают чтение и запись атрибута. 
# При записи автоматически приводит тип

'''
class User(Model):
    name = Field(str, required=True)
    age  = Field(int, default=0)

# ModelMeta собирает _fields при объявлении класса
# ↓
alice = User.create(name="Alice", age=25)
# создаёт User, пишет в data/User.json
# ↓
alice.age = 26
alice.save()
# обновляет запись в файле
# ↓
User.filter(age=26)
# читает файл, фильтрует, возвращает список объектов
'''

class Field:
    def __init__(self, field_type, required=False, default=None):
        self.field_type = field_type
        self.required = required
        self.default = default
        self.name = None  # устанавливается метаклассом

    def __set_name__(self, owner, name):
        self.name = name

    def __get__(self, obj, objtype=None):
        if obj is None:
            return self
        return obj.__dict__.get(self.name, self.default)

    def __set__(self, obj, value):
        if value is not None and not isinstance(value, self.field_type):
            value = self.field_type(value)
        obj.__dict__[self.name] = value

# При создании любого класса-наследника сканирует его тело и собирает все Field-атрибуты в словарь _fields. 
# Это происходит один раз при объявлении класса:
class ModelMeta(type):
    def __new__(mcs, name, bases, namespace):
        fields = {}
        for key, val in namespace.items():
            if isinstance(val, Field):
                fields[key] = val
        namespace["_fields"] = fields
        return super().__new__(mcs, name, bases, namespace)

class Model(metaclass=ModelMeta):
    _fields: dict[str, Field]

    def __init__(self, **kwargs):
        self.id: int = kwargs.pop("id", None)
        for name, field in self._fields.items():
            value = kwargs.get(name, field.default)
            if field.required and value is None:
                raise ValueError(f"Поле '{name}' обязательно")
            setattr(self, name, value)

    @classmethod
    def _db_path(cls) -> str:
        os.makedirs("data", exist_ok=True)
        return f"data/{cls.__name__}.json"

    @classmethod
    def _load_db(cls) -> list[dict]:
        path = cls._db_path()
        if not os.path.exists(path):
            return []
        with open(path, "r", encoding="utf-8") as f:
            return json.load(f)

    @classmethod
    def _save_db(cls, records: list[dict]):
        with open(cls._db_path(), "w", encoding="utf-8") as f:
            json.dump(records, f, ensure_ascii=False, indent=2)

    def _to_dict(self) -> dict:
        d = {"id": self.id}
        for name in self._fields:
            d[name] = getattr(self, name)
        return d

    @classmethod
    def _from_dict(cls, data: dict) -> "Model":
        return cls(**data)

    @classmethod
    def create(cls, **kwargs) -> "Model":
        records = cls._load_db()
        new_id = max((r["id"] for r in records), default=0) + 1
        obj = cls(id=new_id, **kwargs)
        records.append(obj._to_dict())
        cls._save_db(records)
        return obj

    def save(self):
        records = self.__class__._load_db()
        for i, r in enumerate(records):
            if r["id"] == self.id:
                records[i] = self._to_dict()
                break
        self.__class__._save_db(records)

    def delete(self):
        records = self.__class__._load_db()
        records = [r for r in records if r["id"] != self.id]
        self.__class__._save_db(records)

    @classmethod
    def all(cls) -> list["Model"]:
        return [cls._from_dict(r) for r in cls._load_db()]

    @classmethod
    def get(cls, id: int) -> "Model | None":
        for r in cls._load_db():
            if r["id"] == id:
                return cls._from_dict(r)
        return None

    @classmethod
    def filter(cls, **kwargs) -> list["Model"]:
        return [
            cls._from_dict(r) for r in cls._load_db()
            if all(r.get(k) == v for k, v in kwargs.items())
        ]

def main():
    # Задача 1
    print("=== Matrix ===")
    A = Matrix([[1, 2, 3], [4, 5, 6]])
    B = Matrix([[7, 8, 9], [10, 11, 12]])
    C = Matrix([[1, 2], [3, 4], [5, 6]])

    print(A + B)
    print(A * 2)
    print(A.transpose())
    print(A.rows, A.cols)
    print(A[0][2])
    print(A * C)

    # Задача 2
    print("\n=== Zoo ===")
    zoo = Zoo("Московский зоопарк")
    zoo.add(Lion("Симба", 5))
    zoo.add(Dolphin("Флиппер", 8))
    zoo.add(Eagle("Орлан", 3))

    print(zoo)
    for animal in zoo:
        print(f"  {animal.info()}")
        print(f"  -> {animal.speak()}, {animal.move()}")

    print(f"Хищники: {[a.name for a in zoo.by_diet('хищник')]}")
    print(f"Старейший: {zoo.oldest().name}")

    # Задача 3
    print("\n=== Cache TTL ===")

    @cache(ttl=5)
    def slow_fibonacci(n):
        if n <= 1:
            return n
        time.sleep(0.01)
        return slow_fibonacci(n - 1) + slow_fibonacci(n - 2)

    start = time.time()
    print(slow_fibonacci(20))
    print(f"Первый вызов: {time.time() - start:.2f} сек")

    start = time.time()
    print(slow_fibonacci(20))
    print(f"Второй вызов: {time.time() - start:.4f} сек")

    print(slow_fibonacci.cache_info())

    # Задача 4
    print("\n=== Mini ORM ===")

    class User(Model):
        name = Field(str, required=True)
        email = Field(str, required=True)
        age = Field(int, default=0)

    class Post(Model):
        title = Field(str, required=True)
        content = Field(str, default="")
        author_id = Field(int, required=True)

    u1 = User.create(name="Алиса", email="alice@example.com", age=30)
    u2 = User.create(name="Боб", email="bob@example.com", age=25)

    print(u1.id, u1.name)

    u1.age = 31
    u1.save()

    print([u.name for u in User.all()])
    print(User.get(id=1).name)
    print([u.name for u in User.filter(age=25)])

    u2.delete()
    print([u.name for u in User.all()])


if __name__ == "__main__":
    main()
